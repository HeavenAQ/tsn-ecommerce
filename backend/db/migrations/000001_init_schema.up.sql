CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enum Types
CREATE TYPE product_status AS ENUM (
    'in-stock',
    'out-of-stock',
    'discontinued'
);

CREATE TYPE language_code AS ENUM (
    'chn',
    'jp'
);

CREATE TYPE order_status AS ENUM (
    'pending',
    'processing',
    'shipped',
    'delivered',
    'cancelled'
);

-- Languages table
CREATE TABLE "languages" (
    "pk" bigserial PRIMARY KEY,
    "code" language_code NOT NULL DEFAULT 'chn',
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

INSERT INTO languages (code)
    VALUES ('chn'),
    ('jp');

-- Products table
CREATE TABLE "products" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid DEFAULT uuid_generate_v4 (),
    "price" int NOT NULL DEFAULT 0,
    "discount" int NOT NULL DEFAULT 0,
    "image" text NOT NULL DEFAULT '',
    "status" product_status NOT NULL DEFAULT 'in-stock',
    "quantity" int NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "product_translations" (
    "pk" bigserial PRIMARY KEY,
    "product_pk" bigint NOT NULL REFERENCES products (pk) ON DELETE CASCADE,
    "language_pk" bigint NOT NULL REFERENCES languages (pk) ON DELETE CASCADE,
    "name" text NOT NULL DEFAULT '',
    "description" text NOT NULL DEFAULT '',
    "category" text NOT NULL DEFAULT '',
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Users table
CREATE TABLE "users" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid DEFAULT uuid_generate_v4 (),
    "email" text UNIQUE NOT NULL DEFAULT '',
    "phone" text UNIQUE NOT NULL DEFAULT '',
    "password" text NOT NULL DEFAULT '',
    "first_name" text NOT NULL DEFAULT '',
    "last_name" text NOT NULL DEFAULT '',
    "language_pk" bigint NOT NULL REFERENCES languages (pk),
    "address" text NOT NULL DEFAULT '',
    "last_login" timestamptz NOT NULL DEFAULT now(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Orders table
CREATE TABLE "orders" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid DEFAULT uuid_generate_v4 (),
    "user_pk" bigint NOT NULL REFERENCES users (pk) ON DELETE CASCADE,
    "status" order_status NOT NULL DEFAULT 'pending',
    "is_paid" boolean NOT NULL DEFAULT FALSE,
    "total_price" int NOT NULL DEFAULT 0,
    "shipping_address" text NOT NULL,
    "shipping_date" timestamptz,
    "delivered_date" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "order_details" (
    "pk" bigserial PRIMARY KEY,
    "order_pk" bigint NOT NULL REFERENCES orders (pk) ON DELETE CASCADE,
    "product_pk" bigint NOT NULL REFERENCES products (pk) ON DELETE CASCADE,
    "quantity" int NOT NULL DEFAULT 0,
    "price" int NOT NULL DEFAULT 0,
    "discount" int NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE FUNCTION update_updated_at ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

CREATE FUNCTION update_last_login ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.last_login = now();
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER on_update_last_login
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_last_login ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON languages
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON product_translations
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();

CREATE TRIGGER on_update_updated_at
    BEFORE UPDATE ON order_details
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at ();
