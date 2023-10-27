package bootstrap

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/closer"
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

	go func() {
		t := time.NewTicker(cfg.MngPingInterval)
		for range t.C {
			if err = client.Ping(ctx, nil); err != nil {
				log.Error().Msgf("Error pinging mongo, err: %v", err)
			}
		}
	}()

	closer.Add(func() error {
		return client.Disconnect(context.TODO())
	})

	return client
}
