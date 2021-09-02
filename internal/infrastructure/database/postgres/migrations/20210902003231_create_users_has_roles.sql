-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_has_roles (
  users_id uuid NOT NULL,
  roles_id uuid NOT NULL,
  PRIMARY KEY (users_id, roles_id),
  CONSTRAINT fk_users_has_roles_users
    FOREIGN KEY (users_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_users_has_roles_roles1
    FOREIGN KEY (roles_id)
    REFERENCES roles (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

  CREATE INDEX fk_users_has_roles_roles1_idx ON users_has_roles(roles_id ASC);
  CREATE INDEX fk_users_has_roles_users_idx ON users_has_roles(users_id ASC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_has_roles;
-- +goose StatementEnd