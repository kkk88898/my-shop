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
		GRPC         `yaml:"grpc"`
		configs.Log  `yaml:"logger"`
	}

	GRPC struct {
		UserHost    string `env-required:"true" yaml:"user_host" env:"GRPC_USER_HOST"`
		UserPort    int    `env-required:"true" yaml:"user_port" env:"GRPC_USER_PORT"`
		CounterHost string `env-required:"true" yaml:"counter_host" env:"GRPC_COUNTER_HOST"`
		CounterPort int    `env-required:"true" yaml:"counter_port" env:"GRPC_COUNTER_PORT"`
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
