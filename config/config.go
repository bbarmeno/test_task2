package config

import (
	"acc_balance/storage"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address string `env:"ADDRESS,required" envDefault:":8080"`
	SqlCfg  storage.PostgresConfig
}

func (c Config) SqlConfig() storage.PostgresConfig {
	return c.SqlCfg
}

func SetupConfig() (*Config, error) {

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
