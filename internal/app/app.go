package app

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http/handlers"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/bootstrap"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/hash"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/utils"
	"github.com/Kokkibegushidoktor/task-dispenser-service/pkg/storage"
)

func Run(ctx context.Context, cfg *config.Config) error {
	mng := bootstrap.NewMongoClient(ctx, cfg)
	db := mng.Database(cfg.MngDbName)

	fs := bootstrap.NewMinioClient(ctx, cfg)
	storageProvider := storage.NewFileStorage(fs, cfg.FsBucket, cfg.FsEndpoint)

	repos := repository.NewRepositories(db)

	tokenManager, err := auth.NewManager(cfg.JwtSecret)
	if err != nil {
		return err
	}

	hasher := hash.NewSHA1Hasher(cfg.PassSalt)

	services := service.NewServices(service.Deps{
		Repos:           repos,
		TokenManager:    tokenManager,
		Hasher:          hasher,
		AccessTokenTTL:  cfg.AccessTTL,
		StorageProvider: storageProvider,
	})

	hands := handlers.New(services, tokenManager)

	httpServer := http.New(cfg, hands)
	httpServer.Start()

	utils.GracefulShutdown()

	return nil
}
