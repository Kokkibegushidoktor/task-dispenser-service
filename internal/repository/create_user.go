package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/adapters"
)

func (r *RepoImpl) CreateUser(ctx context.Context, req *api.CreateUserRequest) (string, error) {
	dbReq := adapters.CreateUserRequestFromAPI(req)
	user := &models.User{
		Username: dbReq.Username,
		Password: "",
		Admin:    false,
	}

	res, err := r.UsersCl.InsertOne(ctx, &user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", errors.New("user already exists")
		}
		return "", fmt.Errorf("error inserting user to db, err: %v", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error obtaining user id")
	}

	return id.Hex(), nil
}
