package repository

import (
	"context"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
}
