package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/hash"
	"time"
)

type UsersService struct {
	repo           repository.Users
	tokenManager   auth.TokenManager
	hasher         hash.PasswordHasher
	accessTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager, hasher hash.PasswordHasher, accessTTl time.Duration) *UsersService {
	return &UsersService{
		repo:           repo,
		tokenManager:   tokenManager,
		hasher:         hasher,
		accessTokenTTL: accessTTl,
	}
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (string, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetByCredentials(ctx, input.Username, passwordHash)
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
