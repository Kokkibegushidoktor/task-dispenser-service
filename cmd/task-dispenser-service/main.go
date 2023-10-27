package main

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app"
	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal().Msgf("Failed to load environment, err: %v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal().Msgf("Error running service, err: %v", err)
	}
}
