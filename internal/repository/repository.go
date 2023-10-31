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

type Tasks interface {
	Create(ctx context.Context, task *models.Task) (primitive.ObjectID, error)
	GetById(ctx context.Context, taskId primitive.ObjectID) (*models.Task, error)
	Delete(ctx context.Context, taskId primitive.ObjectID) error
}

type Repositories struct {
	Users Users
	Tasks Tasks
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
		Tasks: NewTasksRepo(db),
	}
}
