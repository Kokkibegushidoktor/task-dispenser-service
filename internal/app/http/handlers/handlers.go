package handlers

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/pkg/auth"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
)

type Handlers struct {
	repo         repository.Repository
	tokenManager auth.TokenManager
}

func New(repo repository.Repository, tokenManager auth.TokenManager) *Handlers {
	return &Handlers{
		repo:         repo,
		tokenManager: tokenManager,
	}
}
