package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ApplicationPort int  `env:"APPLICATION_PORT"`
	DebugMode       bool `env:"DEBUG" envDefault:"false"`

	MySqlHost         string `env:"MYSQL_DATABASE_HOST"`
	MySqlDbPort       string `env:"MYSQL_DATABASE_PORT"`
	MySqlUsername     string `env:"MYSQL_DATABASE_USERNAME"`
	MySqlDbPassword   string `env:"MYSQL_DATABASE_PASSWORD"`
	MySqlDatabaseName string `env:"MYSQL_DATABASE_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
