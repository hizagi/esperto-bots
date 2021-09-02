-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
  id uuid DEFAULT uuid_generate_v4(),
  resource_name VARCHAR(45) NOT NULL,
  can_create boolean NOT NULL,
  can_read boolean NOT NULL,
  can_update boolean NOT NULL,
  can_delete boolean NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  PRIMARY KEY (id));

CREATE INDEX resource_name_idx ON permissions(resource_name ASC)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd