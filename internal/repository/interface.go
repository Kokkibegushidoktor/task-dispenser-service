package repository

import (
	"context"
	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *api.CreateUserRequest) (string, error)
	UserLogin(ctx context.Context, req *api.UserLoginRequest) (*api.User, error)
}
