CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id uuid DEFAULT uuid_generate_v4(),
  email VARCHAR(45) NOT NULL,
  password VARCHAR(45) NOT NULL,
  name VARCHAR(45) NOT NULL,
  lastname VARCHAR(45) NOT NULL,
  document VARCHAR(45) NOT NULL,
  birth_date DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NULL,
  deleted_at TIMESTAMPTZ NULL,
  PRIMARY KEY (id));

CREATE INDEX fk_users_email1_idx ON users(email ASC);

insert into users (id, email, password, name, lastname, document, birth_date, created_at) VALUES
('c2d29867-3d0b-d497-9191-18a9d8ee7830', 'joao@joao.com.br', '12456', 'joao', 'ferraz', '1245', '04/02/1232', '2016-06-22 19:10:25-07');



