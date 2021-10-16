package seeds

const SEED_USERS = "SeedUsers"

func (seeder Seeder) SeedUsers() {
	seeder.db.Exec(`insert into users (id, email, password, name, lastname, document, birth_date, created_at) VALUES
		('c2d29867-3d0b-d497-9191-18a9d8ee7830', 'joao@joao.com.br', '12456', 'joao', 'ferraz', '1245', '04/02/1232', '2016-06-22 19:10:25-07');`)
}
