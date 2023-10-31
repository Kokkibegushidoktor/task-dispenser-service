package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSignInInput struct {
	Username string
	Password string
}

type CreateUserInput struct {
	Username string
}

type Users interface {
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
	Create(ctx context.Context, input CreateUserInput) error
}

type CreateTaskInput struct {
	Title       string
	Description string
}

type Tasks interface {
	Create(ctx context.Context, inp CreateTaskInput) (primitive.ObjectID, error)
}

type Services struct {
	Users Users
	Tasks Tasks
}

type Deps struct {
	Repos          *repository.Repositories
	TokenManager   auth.TokenManager
	AccessTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.TokenManager, deps.AccessTokenTTL)
	tasksService := NewTasksService(deps.Repos.Tasks)
	return &Services{
		Users: usersService,
		Tasks: tasksService,
	}
}
