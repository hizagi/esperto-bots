package seeds

import "github.com/kristijorgji/goseeder"

func Init() {
	goseeder.Register(usersSeeder)
	goseeder.Register(categoriesSeeder)
}
