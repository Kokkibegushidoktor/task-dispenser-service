package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{
		repo: repo,
	}
}

func (s *TasksService) Create(ctx context.Context, inp CreateTaskInput) (primitive.ObjectID, error) {
	task := &models.Task{
		Title:       inp.Title,
		Description: inp.Description,
	}

	return s.repo.Create(ctx, task)
}
