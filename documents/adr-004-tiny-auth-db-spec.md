# ADR-004-SPEC: Спецификация схемы базы данных PostgreSQL ядра безопасности tiny-auth-service, стратегия Soft Delete и архитектура репозиториев

## Статус
Принято (Approved) — Сентябрь 2026 г.

## Контекст
В соответствии с системной архитектурой [ADR-002] и gRPC-контрактами [ADR-002-SPEC] микросервис `tiny-auth-service` инкапсулирует в себе весь контур аутентификации, авторизации и управления доступом (IAM/RBAC). База данных этого сервиса является изолированным доверенным периметром безопасности ИБ.

К проектированию физического слоя СУБД и программных абстракций предъявлены следующие требования:
1.  **Строгий контроль целостности ролевой модели (RBAC):** исключение фантомных связей при отзыве прав или блокировке учетных записей.
2.  **Высокая скорость авторизации (High Load):** выборка пользователя со всей его ролевой матрицей должна выполняться за константное время.
3.  **Изоляция секретов:** хранение чувствительных данных (хэши паролей, криптографические ключи) с использованием безопасных типов данных СУБД.
4.  **Канонический Soft Delete и блокировка:** поддержка бизнес-логики экстренного тормоза (Emergency Brake) без физического уничтожения записей из таблиц для обеспечения исторического аудита ИБ.
5.  **Максимальная автономность (Zero Dependencies):** для работы со слоем СУБД запрещено использование внешних ORM и библиотек-мапперов (включая `sqlx`). Вся работа ведется строго через стандартный пакет **`database/sql`**.
6.  **Стандартизация слоя доступа к данным (Repository Pattern):** архитектура репозиториев должна быть построена на базе унифицированных интерфейсов шаблона проектирования (`BaseCRUDRepository` и `BaseOwnedRepository`).

## Решение

### 1. Физическая схема таблиц (PostgreSQL Schema)
Внутри базы данных `tiny-auth` реализуется реляционная схема контроля доступа на основе ролей (RBAC) «многие-ко-многим», состоящая из трех таблиц:

```sql
-- Таблица учетных записей безопасности
CREATE TABLE users (
    id            VARCHAR(50) PRIMARY KEY,        -- Идентификатор UUIDv7 (String контракт)
    name          VARCHAR(100) NOT NULL UNIQUE,   -- Имя/логин в системе (email или телефон)
    user_type     VARCHAR(50) NOT NULL,           -- Тип учетной записи (e.g., 'b2b', 'system')
    password_hash VARCHAR(255) NOT NULL,          -- Криптографический хэш пароля (argon2id/bcrypt)
    public_key    TEXT NOT NULL DEFAULT '',       -- Асимметричный открытый ключ пользователя
    private_key   TEXT NOT NULL DEFAULT '',       -- Зашифрованный закрытый ключ пользователя
    active        BOOLEAN NOT NULL DEFAULT TRUE,  -- Флаг блокировки (Emergency Brake)
    deleted       BOOLEAN NOT NULL DEFAULT FALSE, -- Флаг Soft Delete (удаление аккаунта)
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Таблица доступных в системе ролей безопасности
CREATE TABLE roles (
    id          VARCHAR(50) PRIMARY KEY,        -- Идентификатор роли UUIDv7 (String контракт)
    name        VARCHAR(50) NOT NULL UNIQUE,    -- Уникальный строковый маркер (e.g., 'admin', 'manager')
    description TEXT NOT NULL DEFAULT '',       -- Описание обязанностей роли
    deleted     BOOLEAN NOT NULL DEFAULT FALSE, -- Флаг Soft Delete
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Связующая таблица RBAC (многие-ко-многим)
CREATE TABLE user_roles (
    user_id     VARCHAR(50) NOT NULL,
    role_id     VARCHAR(50) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles(id)
);
```

### 2. Спецификация Архитектуры Репозиториев (Repository Spec)
Для унификации кодовой базы репозитории микросервиса наследуют интерфейсы и структуру базовых типов шаблона проектирования `go-service-template`:

*   **`RoleRepository`**: Реализует паттерн **`BaseCRUDRepository[string, *domain.Role]`**. Он закрывает стандартный атомарный CRUD (поиск по ID, пагинированный список ролей, идемпотентный апсерт, удаление флагом).
*   **`UserRepository`**: Расширяет шаблон и реализует **`BaseOwnedRepository[string, *domain.User]`** (или кастомный интерфейс поверх него). Это необходимо, так как помимо базового CRUD для управления пользователем, репозиторий инкапсулирует в себе атомарную синхронизацию каскадной коллекции `User.Roles` внутри единой транзакции базы данных.

### 3. Производительность и Индексация (Indexes)
Для предотвращения деградации производительности СУБД при операциях поиска по имени, а также для оптимизации каскадных проверок и удаления ролей настраиваются следующие индексы:

```sql
-- Ускорение поиска по логину (метод FindByName и AuthService.Login)
CREATE INDEX IF NOT EXISTS idx_users_name_active ON users(name) WHERE (deleted = FALSE);

-- Инвертированный индекс связующей таблицы. Составной PK закрывает поиск по (user_id, role_id).
-- Индекс по (role_id, user_id) важен для каскадных проверок связей на уровне Base-шаблона.
CREATE INDEX IF NOT EXISTS idx_user_roles_inverse ON user_roles(role_id ASC, user_id ASC);
```

### 4. Паттерн выборки данных (нативный `database/sql`)
Поскольку на уровне проекта принято решение отказаться от автоматических библиотек-мапперов, агрегация связанных сущностей ролей внутрь структуры пользователя (1-ко-многим) реализуется внутри `UserRepository` вручную через классический паттерн сканирования плоского SQL-ответа:

```go
const fetchUserWithRolesQuery = `
    SELECT 
        u.id, u.name, u.user_type, u.password_hash, u.public_key, u.private_key, u.active, u.deleted, u.created_at, u.updated_at,
        r.id, r.name, r.description, r.deleted, r.created_at, r.updated_at
    FROM users u
    LEFT JOIN user_roles ur ON u.id = ur.user_id
    LEFT JOIN roles r ON ur.role_id = r.id AND r.deleted = FALSE
    WHERE u.name = $1 AND u.deleted = FALSE;
`
```
*Инженерная практика в репозитории:* Внутри цикла `rows.Next()` поля структуры `User` сканируются один раз, а поля `Role` проверяются на `sql.NullString` и динамически добавляются в срез `User.Roles` на каждой итерации. Поля `PublicKey`, `PrivateKey` и `Description` застрахованы на уровне СУБД констрейнтом `NOT NULL DEFAULT`, что гарантирует стабильный скан в плоские строки Go.

### 5. Стратегия дезактивации (Soft Delete & Emergency Brake)
*   **Изоляция блокировок:** Флаг `users.active = FALSE` отвечает за мгновенное срабатывание механизма **Emergency Brake**. Все SQL-запросы на аутентификацию (`AuthService.Login`) в обязательном порядке содержат предикат `WHERE active = TRUE`.
*   **Мягкое удаление (Soft Delete):** При вызове `Delete` на сущностях `User` или `Role` на СУБД накладывается ограничение изменения флага `deleted = TRUE`. Физическое стирание строк `DELETE FROM` запрещено для сохранения истории аудита ИБ (Day-2 Operations).

## Последствия

### Плюсы
*   **Стандартизация кодовой базы:** Использование `BaseCRUDRepository` и `BaseOwnedRepository` гарантирует, что базовые методы работы с базой данных написаны в едином стиле по всему проекту.
*   **Низкий уровень бойлерплейта (Переиспользуемость):** Вопреки ограничениям нативного `database/sql`, лапша из дублирующихся вызовов `rows.Scan()` исключена. Логика маппинга полей описывается ровно **один раз на репозиторий** через выделенную функцию-сканнер, которая переиспользуется как в базовых CRUD-выборках, так и в сложных кастомных `JOIN`-запросах.
*   **Нулевые накладные расходы (Zero Overhead):** Полный отказ от сторонних библиотек гарантирует отсутствие рефлексии (Reflection) в рантайме, обеспечивая максимальную скорость и снижая нагрузку на Garbage Collector.
*   **Атомарность прав (ACID):** Модификация ролей и блокировка учетных записей выполняются внутри нативных ACID-транзакций, исключая рассинхронизацию ИБ контура.

### Минусы / Риски
*   **Синхронизация схемы данных:** Любое изменение схемы таблицы (добавление/удаление колонки в Postgres) по-прежнему требует точечной ручной правки функции-сканнера в Go-коде, так как автоматический маппинг по тегам структур отсутствует.
