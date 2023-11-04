package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TasksService struct {
	repo          repository.Tasks
	levelsService Levels
}

func NewTasksService(repo repository.Tasks, levelsService Levels) *TasksService {
	return &TasksService{
		repo:          repo,
		levelsService: levelsService,
	}
}

func (s *TasksService) Create(ctx context.Context, inp CreateTaskInput) (primitive.ObjectID, error) {
	task := &models.Task{
		Title:       inp.Title,
		Description: inp.Description,
	}

	return s.repo.Create(ctx, task)
}

func (s *TasksService) Update(ctx context.Context, inp UpdateTaskInput) error {
	id, err := primitive.ObjectIDFromHex(inp.ID)
	if err != nil {
		return err
	}

	repoInput := repository.UpdateTaskInput{
		ID:          id,
		Title:       inp.Title,
		Description: inp.Description,
	}

	return s.repo.Update(ctx, repoInput)
}

func (s *TasksService) Delete(ctx context.Context, taskId string) error {
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return s.levelsService.DeleteByTaskId(ctx, id)
}
