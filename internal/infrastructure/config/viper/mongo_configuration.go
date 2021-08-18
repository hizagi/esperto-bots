package viper

import "fmt"

type MongoConfiguration struct {
	DatabaseConfiguration
}

func (configuration *MongoConfiguration) GetDatabaseURL() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/",
		configuration.Username,
		configuration.Password,
		configuration.Host,
		configuration.Port,
	)
}
