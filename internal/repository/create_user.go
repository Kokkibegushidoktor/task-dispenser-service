package repository

import (
	"context"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
)

func (r *RepoImpl) CreateUser(ctx context.Context, user *models.User) (string, error) {
	return "nil", nil
}
