package repository

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user *models.User) error
	GetByCredentials(ctx context.Context, username, password string) (*models.User, error)
}

type UpdateTaskInput struct {
	ID    primitive.ObjectID
	Title string
}

type Tasks interface {
	Create(ctx context.Context, task *models.Task) (primitive.ObjectID, error)
	Update(ctx context.Context, inp UpdateTaskInput) error
	GetById(ctx context.Context, taskId primitive.ObjectID) (*models.Task, error)
	Delete(ctx context.Context, taskId primitive.ObjectID) error
}

type UpdateLevelInput struct {
	ID            primitive.ObjectID
	Title         string
	VarQuestCount int
}

type Levels interface {
	Create(ctx context.Context, level *models.TaskLevel) (primitive.ObjectID, error)
	Update(ctx context.Context, inp UpdateLevelInput) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetByTaskId(ctx context.Context, id primitive.ObjectID) ([]models.TaskLevel, error)
	AddQuestion(ctx context.Context, id primitive.ObjectID, question *models.LevelQuestion) error
	UpdateQuestion(ctx context.Context, inp *models.LevelQuestion) error
	DeleteQuestion(ctx context.Context, id primitive.ObjectID) error
}

type Repositories struct {
	Users  Users
	Tasks  Tasks
	Levels Levels
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:  NewUsersRepo(db),
		Tasks:  NewTasksRepo(db),
		Levels: NewLevelsRepo(db),
	}
}
