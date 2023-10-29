package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
)

type RepoImpl struct {
	client  *mongo.Client
	UsersCl *mongo.Collection
}

func New(cfg *config.Config, client *mongo.Client) *RepoImpl {
	return &RepoImpl{
		client:  client,
		UsersCl: client.Database(cfg.MngDbName).Collection("users"),
	}
}
