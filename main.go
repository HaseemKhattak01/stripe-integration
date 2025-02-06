package main

import (
	"log"

	"github.com/HaseemKhattak01/stripe-integration/app"
	"github.com/HaseemKhattak01/stripe-integration/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err = cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	app.StartServer(cfg) // Starting the server
}
