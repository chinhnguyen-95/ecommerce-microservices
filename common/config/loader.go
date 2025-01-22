package config

import (
	"log"

	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
)

// LoadConfig loads the shared configuration from environment variables.
func LoadConfig() (*AppConfig, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	err = conf.Load(env.NewSource())
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	if err := conf.Scan(&appConfig); err != nil {
		return nil, err
	}

	log.Printf("Configuration loaded: %+v", appConfig)
	return &appConfig, nil
}
