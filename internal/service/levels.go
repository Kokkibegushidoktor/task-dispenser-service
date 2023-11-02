package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LevelsService struct {
	repo repository.Levels
}

func NeLevelsService(repo repository.Levels) *LevelsService {
	return &LevelsService{
		repo: repo,
	}
}

func (s *LevelsService) Create(ctx context.Context, inp CreateLevelInput) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(inp.TaskId)
	if err != nil {
		return id, err
	}

	level := &models.TaskLevel{
		Title:         inp.Title,
		TaskId:        id,
		VarQuestCount: inp.VarQuestCount,
	}

	return s.repo.Create(ctx, level)
}

func (s *LevelsService) Update(ctx context.Context, inp UpdateLevelInput) error {
	id, err := primitive.ObjectIDFromHex(inp.ID)
	if err != nil {
		return err
	}

	repoInput := repository.UpdateLevelInput{
		ID:            id,
		Title:         inp.Title,
		VarQuestCount: inp.VarQuestCount,
	}

	return s.repo.Update(ctx, repoInput)
}
