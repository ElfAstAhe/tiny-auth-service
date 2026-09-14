-- Manual usage

-- 1. Позволяет создавать новые объекты (таблицы, индексы, последовательности)
GRANT CREATE ON SCHEMA auth_db TO svc_auth;

-- 2. Позволяет видеть схему
GRANT USAGE ON SCHEMA auth_db TO svc_auth;

-- 3. Если таблицы уже созданы другим юзером (например, postgres),
-- нужно сделать test их владельцем или дать полные права:
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA auth_db TO svc_auth;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA auth_db TO svc_auth;

-- 4. Чтобы будущие объекты тоже были под контролем:
ALTER DEFAULT PRIVILEGES IN SCHEMA auth_db
    GRANT ALL PRIVILEGES ON TABLES TO svc_auth;
ALTER DEFAULT PRIVILEGES IN SCHEMA auth_db
    GRANT ALL PRIVILEGES ON SEQUENCES TO svc_auth;
