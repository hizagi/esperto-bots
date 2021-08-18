package viper

type Configuration struct {
	MongoConfiguration MongoConfiguration `mapstructure:"mongo_configuration"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
