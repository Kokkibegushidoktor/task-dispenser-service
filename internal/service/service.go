package service

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/hash"
	"github.com/Kokkibegushidoktor/task-dispenser-service/pkg/storage"
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

type UserSetUpPassInput struct {
	Username string
	Password string
}

type Users interface {
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
	SetUpPassword(ctx context.Context, inp UserSetUpPassInput) error
	Create(ctx context.Context, input CreateUserInput) error
}

type CreateTaskInput struct {
	Title       string
	Description string
}

type UpdateTaskInput struct {
	ID          string
	Title       string
	Description string
}

type Tasks interface {
	Create(ctx context.Context, inp CreateTaskInput) (primitive.ObjectID, error)
	Update(ctx context.Context, inp UpdateTaskInput) error
	Delete(ctx context.Context, taskId string) error
}

type CreateLevelInput struct {
	Title         string
	VarQuestCount int
	TaskId        string
}

type UpdateLevelInput struct {
	ID            string
	Title         string
	VarQuestCount int
}

type Levels interface {
	Create(ctx context.Context, inp CreateLevelInput) (primitive.ObjectID, error)
	Update(ctx context.Context, inp UpdateLevelInput) error
	Delete(ctx context.Context, levelId string) error
	DeleteByTaskId(ctx context.Context, taskId primitive.ObjectID) error
}

type AddQuestionInput struct {
	LevelID     string
	Title       string
	Description string
}

type UpdateQuestionInput struct {
	ID          string
	Title       string
	Description string
}

type Questions interface {
	Create(ctx context.Context, inp AddQuestionInput) (primitive.ObjectID, error)
	Update(ctx context.Context, inp UpdateQuestionInput) error
	Delete(ctx context.Context, id string) error
}

type Files interface {
	SaveAndUpload(ctx context.Context, file *models.File) (string, error)
}

type Services struct {
	Users     Users
	Tasks     Tasks
	Levels    Levels
	Questions Questions
	Files     Files
}

type Deps struct {
	Repos           *repository.Repositories
	TokenManager    auth.TokenManager
	Hasher          hash.PasswordHasher
	AccessTokenTTL  time.Duration
	StorageProvider storage.Provider
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.TokenManager, deps.Hasher, deps.AccessTokenTTL)
	levelsService := NewLevelsService(deps.Repos.Levels)
	tasksService := NewTasksService(deps.Repos.Tasks, levelsService)
	questionsService := NewQuestionsService(deps.Repos.Levels)
	fileService := NewFileService(deps.StorageProvider, deps.Repos.Files)
	return &Services{
		Users:     usersService,
		Tasks:     tasksService,
		Levels:    levelsService,
		Questions: questionsService,
		Files:     fileService,
	}
}
