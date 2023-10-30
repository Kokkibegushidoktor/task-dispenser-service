package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"time"
)

type UsersService struct {
	repo           repository.Users
	tokenManager   auth.TokenManager
	accessTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager, accessTTl time.Duration) *UsersService {
	return &UsersService{
		repo:           repo,
		tokenManager:   tokenManager,
		accessTokenTTL: accessTTl,
	}
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (string, error) {
	user, err := s.repo.GetByCredentials(ctx, input.Username, input.Password)
	if err != nil {
		return "", err
	}

	token, err := s.tokenManager.NewJWT(user, s.accessTokenTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UsersService) Create(ctx context.Context, input CreateUserInput) error {
	user := &models.User{
		Username: input.Username,
		Password: "",
		Admin:    false,
	}
	err := s.repo.Create(ctx, user)

	return err
}
