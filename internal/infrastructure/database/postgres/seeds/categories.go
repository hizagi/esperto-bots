package seeds

import (
	"github.com/kristijorgji/goseeder"
)

func categoriesSeeder(s goseeder.Seeder) {
	s.DB.Exec(`insert into categories (id, name, slug, parent_id, created_at) VALUES
		('c2d29867-3d0b-d497-9191-18a9d8ee7830', 'Vacuum', 'vacuum', null, '2016-06-20 19:10:25-07'),
		('2a3eb41e-2ccf-4357-a7cb-b0d491a53b96', 'Power Vacuum', 'power-vacuum', 'c2d29867-3d0b-d497-9191-18a9d8ee7830', '2016-06-21 19:10:25-07'),
		('87169909-33c1-4a2c-8c77-76cba19369f1', 'Mop', 'mop', null, '2016-06-22 19:10:25-07'),
		('e874153d-7aab-4fa9-bb5a-21082bf7ced6', 'Power Mop', 'power-mop', '87169909-33c1-4a2c-8c77-76cba19369f1', '2016-06-23 19:10:25.001-07');`)
}
