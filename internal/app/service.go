package app

import (
	"context"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/bootstrap"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
)

func Run(ctx context.Context, cfg *config.Config) error {
	mng := bootstrap.New(ctx, cfg)
	repo := repository.New(cfg, mng)

	httpServer := http.New(cfg, repo)
	httpServer.Start()

	return nil
}
