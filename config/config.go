package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	BindPort string `env:"BIND_PORT"envDefault:":8080"`
	DSN      string `env:"DB_CONN,required"`
	DBPwd    string `env:"DB_PWD"envDefault:""`
	FileName string `env:"FILE_NAME"envDefault:"data-20190906T0100.json"`
}

func Get() (*Config, error) {
	conf := Config{}
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
