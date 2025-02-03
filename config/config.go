package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StripeKey string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	stripeKey := os.Getenv("STRIPE_API_KEY")
	if stripeKey == "" {
		return nil, fmt.Errorf("STRIPE_API_KEY environment variable not set")
	}

	return &Config{StripeKey: stripeKey}, nil
}

func (c *Config) Validate() error {
	if c.StripeKey == "" {
		return fmt.Errorf("StripeKey is not set in the configuration")
	}
	return nil
}