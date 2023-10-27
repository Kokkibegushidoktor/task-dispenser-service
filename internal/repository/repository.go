package repository

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoImpl struct {
	client *mongo.Client
}

func New(cfg *config.Config, client *mongo.Client) *RepoImpl {
	return &RepoImpl{
		client: client,
	}
}
