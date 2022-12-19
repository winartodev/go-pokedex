package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Application struct {
		Environment string `env:"APP_ENV,required"`
		Host        string `env:"APP_URL,required"`
		Port        int64  `env:"APP_PORT,required"`
	}

	Database struct {
		Connection string `env:"DB_CONNECTION,required"`
		Username   string `env:"DB_USERNAME,required"`
		Password   string `env:"DB_PASSWORD,required"`
		Host       string `env:"DB_HOST,default=localhost"`
		Port       string `env:"DB_PORT,required"`
		Database   string `env:"DB_DATABASE,required"`
	}
}

// NewConfig will return the Config read from the .env file
func NewConfig() Config {
	var cfg Config
	gotenv.Load(".env")
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
