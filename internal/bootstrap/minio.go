package bootstrap

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func NewMinioClient(ctx context.Context, cfg *config.Config) *minio.Client {
	client, err := minio.New(cfg.FsEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.FsAccessKey, cfg.FsSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatal().Msgf("Error connecting to endpoint, err: %v", err)
	}

	exists, err := client.BucketExists(ctx, cfg.FsBucket)
	if err != nil {
		log.Fatal().Msgf("Error checking specified bucket, err: %v", err)
	}
	if !exists {
		log.Fatal().Msg("Specified bucket does not exist")
	}

	return client
}
