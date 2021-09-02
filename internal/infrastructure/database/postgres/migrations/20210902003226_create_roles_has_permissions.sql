-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles_has_permissions (
  roles_id uuid NOT NULL,
  permissions_id uuid NOT NULL,
  PRIMARY KEY (permissions_id, roles_id),
  CONSTRAINT fk_roles_has_permissions_permissions
    FOREIGN KEY (permissions_id)
    REFERENCES permissions (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_roles_has_permissions_roles1
    FOREIGN KEY (roles_id)
    REFERENCES roles (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

  CREATE INDEX fk_roles_has_permissions_roles1_idx ON roles_has_permissions(roles_id ASC);
  CREATE INDEX fk_roles_has_permissions_permissions_idx ON roles_has_permissions(permissions_id ASC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles_has_permissions;
-- +goose StatementEnd