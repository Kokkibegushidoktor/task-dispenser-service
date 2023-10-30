package handlers

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
)

type Handlers struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func New(services *service.Services, tokenManager auth.TokenManager) *Handlers {
	return &Handlers{
		services:     services,
		tokenManager: tokenManager,
	}
}
