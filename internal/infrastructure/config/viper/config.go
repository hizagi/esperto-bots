package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	MongoConfiguration MongoConfiguration `mapstructure:"mongo_configuration"`
}

func NewConfiguration(rootPath string) *Configuration {
	viper.SetConfigName("env")    // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(rootPath) // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	fmt.Printf("Testing Config: %+v", viper.AllKeys())

	var configuration Configuration

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return &Configuration{}
	}

	return &configuration
}
