CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE categories (
	id serial,
	"name" varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);

CREATE TABLE  products(
	id serial ,
    category_id int NOT NULL,
	"name" varchar NOT NULL,
    price decimal(10,2),
    stock smallint,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

ALTER TABLE ONLY products ADD CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL