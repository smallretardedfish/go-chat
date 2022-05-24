package configs

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"net/url"
)

type Config struct {
	ServerAddress string `env:"HTTP_SERVER_ADDRESS" envDefault:"http://0.0.0.0:8080"`
	DSN           string `env:"DB_DSN" envDefault:"host=0.0.0.0 user=postgres password=postgres dbname=chat port=5432 sslmode=disable"`
}

func (c *Config) GetServerAddress() (*url.URL, error) {
	return url.Parse(c.ServerAddress)
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
