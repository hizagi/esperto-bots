-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
  id uuid DEFAULT uuid_generate_v4(),
  name VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  PRIMARY KEY (id));

CREATE INDEX name_idx ON roles(name ASC)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd