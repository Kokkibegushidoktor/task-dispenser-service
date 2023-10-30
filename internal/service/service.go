package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"time"
)

type UserSignInInput struct {
	Username string
	Password string
}

type CreateUserInput struct {
	Username string
}

type Users interface {
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
	Create(ctx context.Context, input CreateUserInput) error
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos          *repository.Repositories
	TokenManager   auth.TokenManager
	AccessTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.TokenManager, deps.AccessTokenTTL)
	return &Services{
		Users: usersService,
	}
}
