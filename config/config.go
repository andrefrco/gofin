package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env/v6"
)

// Config of the app
type Config struct {
	Port        string `env:"PORT" envDefault:"3000"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://localhost:5432/gofin?sslmode=disable"`
}

// Get the config
func Get() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return
}
