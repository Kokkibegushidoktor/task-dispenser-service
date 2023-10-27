package bootstrap

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
)

func New(ctx context.Context, cfg *config.Config) *mongo.Client {
	opts := options.Client().
		ApplyURI(cfg.MngDsn).
		SetMaxPoolSize(uint64(cfg.MngMaxPoolSize)).
		SetMinPoolSize(uint64(cfg.MngMinPoolSize)).
		SetMaxConnecting(uint64(cfg.MngMaxConnecting))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal().Msgf("Error connecting to mongodb, err: %v", err)
	}

	return client
}
