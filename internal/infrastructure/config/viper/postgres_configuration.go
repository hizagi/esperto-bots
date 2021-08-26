package viper

import (
	"fmt"
)

type PostgresConfiguration struct {
	DatabaseConfiguration `mapstructure:",squash"`
	MaxConnections        int32 `mapstructure:"max_connections"`
	MaxConnectionIdleTime int   `mapstructure:"max_connection_idle_time"`
	MaxConnectionLifetime int   `mapstructure:"max_connection_lifetime"`
}

func (configuration *PostgresConfiguration) GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		configuration.Username,
		configuration.Password,
		configuration.Host,
		configuration.Port,
		configuration.DatabaseName,
	)
}
