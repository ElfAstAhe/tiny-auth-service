# ADR-002-SPEC: Полная gRPC-спецификация контрактов и моделей данных ядра безопасности tiny-auth-service

## Статус
Принято (Approved) — Апрель 2026 г.

## Контекст
В соответствии с системной архитектурой [ADR-002] микросервис `tiny-auth-service` инкапсулирует в себе весь контур аутентификации, авторизации и управления доступом (IAM/RBAC). Для обеспечения полиморфизма интерфейсов, высокой производительности внутрисетевого взаимодействия и Day-2 Operations (эксплуатации) утверждается монолитная gRPC-спецификация, разделенная на три логических слоя API:
1.  **Административный слой (Admin API):** Полный CRUD над пользователями и ролями для системных инженеров и автоматизации.
2.  **Слой аутентификации (Public Auth API):** Публичные точки входа для обмена учетных данных на сессионные токены.
3.  **Слой самообслуживания (User Self-Service API):** Операции управления аккаунтом со стороны авторизованного пользователя.

## Решение

### 1. Единые инженерные стандарты Protobuf
*   **Синтаксис:** Весь пакет контрактов переведен на современный стандарт **`edition = "2023"`**, исключающий устаревшие конструкции proto3.
*   **Пакетное пространство:** Все сервисы изолированы в рамках единого доменного неймсейса `package auth.service`.
*   **Изоляция типов (Go):** Таргет компиляции фиксирует общий путь генерации: `option go_package = "tiny-auth-service/grpc/tiny-auth-service"`.

### 2. Слой моделей данных (v1/messages.proto)
Сущности спроектированы с учетом стратегии Soft Delete (`deleted`) и криптографического контроля доступов (пары асимметричных ключей). Связующая таблица `UserRoles` упразднена на уровне API за счет агрегации ролей прямо внутрь объекта пользователя.

```protobuf
message Role {
  string id = 1;                  // UUID роли
  string name = 2;                // Уникальный строковый маркер (e.g., "admin")
  string description = 3;         // Описание обязанностей роли
  bool deleted = 4;               // Флаг мягкого удаления
  google.protobuf.Timestamp created_at = 5;
  google.protobuf.Timestamp updated_at = 6;
}

message User {
  string id = 1;                  // UUID пользователя
  string name = 2;                // Имя/логин в системе
  string user_type = 3;           // Тип учетной записи (e.g., "b2b", "system")
  string password_hash = 4;       // Криптографический хэш пароля
  string public_key = 5;          // Асимметричный открытый ключ
  string private_key = 6;         // Зашифрованный закрытый ключ
  bool active = 7;                // Флаг блокировки (Emergency Brake)
  bool deleted = 8;               // Флаг мягкого удаления
  google.protobuf.Timestamp created_at = 9;
  google.protobuf.Timestamp updated_at = 10;
  repeated Role roles = 11;       // Агрегированная матрица доступов (RBAC)
}
```

### 3. Карта gRPC-сервисов и методов

#### 3.1 Административный слой (Admin API)
Представлен двумя симметричными сервисами `AdminUsersService` и `AdminRolesService`. Реализует паттерн идемпотентного сохранения (`Save` выполняет роль Upsert-операции) и пагинацию на уровне СУБД (`limit` / `offset`).

*   **`AdminUsersService`**
    *   `rpc Find(AdminUsersFindRequest) returns (AdminUsersInstanceResponse);` (Поиск по UUID)
    *   `rpc FindByName(AdminUsersFindByNameRequest) returns (AdminUsersInstanceResponse);` (Поиск по логину)
    *   `rpc List(AdminUsersListRequest) returns (AdminUsersInstancesResponse);` (Пагинированный список)
    *   `rpc Save(AdminUsersSaveRequest) returns (AdminUsersInstanceResponse);` (Создание/Обновление)
    *   `rpc Delete(AdminUsersDeleteRequest) returns (google.protobuf.Empty);` (Удаление)
*   **`AdminRolesService`**
    *   `rpc Find(AdminRolesFindRequest) returns (AdminRolesInstanceResponse);`
    *   `rpc FindByName(AdminRolesFindByNameRequest) returns (AdminRolesInstanceResponse);`
    *   `rpc List(AdminRolesListRequest) returns (AdminRolesInstancesResponse);`
    *   `rpc Save(AdminRolesSaveRequest) returns (AdminRolesInstanceResponse);`
    *   `rpc Delete(AdminRolesDeleteRequest) returns (google.protobuf.Empty);`

#### 3.2 Слой аутентификации (AuthService)
Отвечает за базовый выпуск сессий. Предоставляет метод классического логина и легковесного (`LoginSimple`) для снижения накладных расходов сети.
```protobuf
service AuthService {
  rpc Login(AuthLoginRequest) returns (AuthLoginResponse);
  rpc LoginSimple(AuthLoginRequest) returns (AuthLoginResponse);
}

message AuthLoginRequest {
  string username = 1;
  string password = 2;
}

message AuthLoginResponse {
  string token = 1;          // Короткоживущий access-токен (JWT)
  string refresh_token = 2;  // Долговременный токен обновления сессии
}
```

#### 3.3 Слой самообслуживания (UserService)
Обеспечивает изоляцию контекста текущей сессии (методы не требуют явной передачи `user_id`, так как он извлекается из контекста метаданных токена gRPC-интерцептором).
```protobuf
service UserService {
  rpc Profile(google.protobuf.Empty) returns (ProfileResponse);
  rpc ChangePassword(ChangePasswordRequest) returns (google.protobuf.Empty);
  rpc ChangeKeys(google.protobuf.Empty) returns (ChangeKeysResponse);
}

message ChangePasswordRequest {
  string old_password = 1;
  string new_password = 2;
}

message ProfileResponse {
  string id = 1;
  string name = 2;
  string user_type = 3;
  string public_key = 4;
  bool active = 5;
  google.protobuf.Timestamp created_at = 6;
  google.protobuf.Timestamp updated_at = 7;
  repeated string roles = 8; // Сжатый строковый формат ролей для оптимизации трафика клиента
}

message ChangeKeysResponse {
  string public_key = 1;
  string private_key = 2;
}
```

## Последствия

### Плюсы
*   **Высокая консистентность (API Design Symmetry):** Идентичные структуры административных методов упрощают генерацию кода и написание сквозных middleware-компонентов.
*   **Оптимизация сетевого трафика (Payload Trimming):** `UserService.Profile` не гонит тяжелые каскадные объекты `Role` (с таймстампами и описаниями), а отдает плоский срез строк `repeated string roles`, снижая нагрузку на парсинг на мобильных клиентах.
*   **Изоляция секретов:** Метод `Profile` на уровне контракта исключает передачу `password_hash` и `private_key` на сторону клиента, предотвращая случайную утечку конфиденциальных данных.

### Минусы / Риски
*   **Head-of-Line Blocking на пагинации:** Паттерн пагинации через `offset` при экстремально больших объемах данных в таблице `users` начнет замедлять SQL-запросы в Postgres. При масштабировании базы до миллионов записей потребуется перевод контракта `List` на Keyset Pagination (Cursor-based).
