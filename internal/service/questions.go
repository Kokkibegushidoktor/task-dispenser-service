package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionsService struct {
	repo repository.Levels
}

func NewQuestionsService(repo repository.Levels) *QuestionsService {
	return &QuestionsService{
		repo: repo,
	}
}

func (s *QuestionsService) Create(ctx context.Context, inp AddQuestionInput) (primitive.ObjectID, error) {
	levelId, err := primitive.ObjectIDFromHex(inp.LevelID)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	question := &models.LevelQuestion{
		ID:          primitive.NewObjectID(),
		Title:       inp.Title,
		Description: inp.Description,
	}

	return question.ID, s.repo.AddQuestion(ctx, levelId, question)
}

func (s *QuestionsService) Update(ctx context.Context, inp UpdateQuestionInput) error {
	id, err := primitive.ObjectIDFromHex(inp.ID)
	if err != nil {
		return err
	}

	repoInput := &models.LevelQuestion{
		ID:          id,
		Title:       inp.Title,
		Description: inp.Description,
		ContentURL:  inp.ContentURL,
	}

	return s.repo.UpdateQuestion(ctx, repoInput)
}

func (s *QuestionsService) Delete(ctx context.Context, qid string) error {
	id, err := primitive.ObjectIDFromHex(qid)
	if err != nil {
		return err
	}
	return s.repo.DeleteQuestion(ctx, id)
}
