-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
  id uuid DEFAULT uuid_generate_v4(),
  parent_id uuid NULL,
  name VARCHAR(45) NOT NULL,
  slug VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NULL,
  deleted_at TIMESTAMPTZ(6) NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_categories_categories1
    FOREIGN KEY (parent_id)
    REFERENCES categories (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);
  CREATE INDEX fk_categories_keyset_idx ON categories(created_at, id);
  CREATE INDEX fk_categories_categories1_idx ON categories(parent_id ASC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
