package handlers

import "github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"

type Handlers struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}
