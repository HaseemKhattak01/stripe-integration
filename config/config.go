package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StripeKey string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	stripeKey := os.Getenv("STRIPE_API_KEY")
	if stripeKey == "" {
		return nil, errors.New("STRIPE_API_KEY environment variable not set")
	}

	return &Config{StripeKey: stripeKey}, nil
}

func (c *Config) Validate() error {
	if c.StripeKey == "" {
		return errors.New("StripeKey is not set in the configuration")
	}
	return nil
}
