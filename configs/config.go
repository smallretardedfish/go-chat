package configs

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"net/url"
)

type Config struct {
	ServerAddress string `env:"HTTP_SERVER_ADDRESS" envDefault:":8080"`
	DSN           string `env:"DB_DSN" envDefault:"host=0.0.0.0 user=postgres password=mysecretpassword dbname=chatdb port=5432 sslmode=disable"`
	jwtKey        string `env:"JWT_SECRET_KEY"`
}

func (c *Config) GetServerAddress() (*url.URL, error) {
	return url.Parse(c.ServerAddress)
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
