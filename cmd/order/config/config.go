package config

import (
	"fmt"
	"log"
	"os"

	configs "myshop/pkg/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
		MYSQL8       `yaml:"mysql8"`
	}

	MYSQL8 struct {
		PoolMax     int    `env-required:"true" yaml:"pool_max" env:"MYSQL8_POOL_MAX"`
		IdleConnMax int    `env-required:"true" yaml:"idle_conn_max" env:"MYSQL8_IDLE_CONN_MAX"`
		MaxIdleTime int    `env-required:"true" yaml:"max_idle_time" env:"MYSQL8_MAX_IDLE_TIME"`
		DsnURL      string `env-required:"true" yaml:"dsn_url" env:"MYSQL8_DSN_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
