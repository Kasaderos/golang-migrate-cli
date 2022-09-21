package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MigrationsURL string         `json:"migrations_url"`
	Postgres      PostgresConfig `json:"postgres"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"posrt"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func GetConfig() (c Config, err error) {
	var filePath string
	var config Config

	if os.Getenv("config") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return c, err
		}

		filePath = pwd + "/internal/app/config/config.json"
	} else {
		filePath = os.Getenv("config")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return c, err
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return c, err
	}

	return config, nil
}
