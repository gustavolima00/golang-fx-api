package config

import (
	"log"

	env "github.com/caarlos0/env/v11"
)

// Config defines the application env vars
type Config struct {
	Port             string `env:"PORT" envDefault:"8080"`
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresDatabase string `env:"POSTGRES_DATABASE" envDefault:"mydb"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"myuser"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"mypassword"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDB       string `env:"POSTGRES_DB" envDefault:"mydb"`
}

// New will parse the necessary env vars to
// struct Config
func New() *Config {
	c := new(Config)

	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	return c
}
