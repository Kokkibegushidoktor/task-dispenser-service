package repository

import (
	"context"
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	col *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		col: db.Collection("users"),
	}
}

func (r *UsersRepo) Create(ctx context.Context, user *models.User) error {
	_, err := r.col.InsertOne(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return models.ErrAlreadyExists
	}

	return err
}

func (r *UsersRepo) Update(ctx context.Context, inp UpdateUserInput) error {
	updateQuery := bson.M{}

	if inp.Username != "" {
		updateQuery["username"] = inp.Username
	}

	if inp.Password != "" {
		updateQuery["password"] = inp.Password
	}

	res, err := r.col.UpdateByID(ctx, inp.ID, bson.M{"$set": updateQuery})
	if res.MatchedCount == 0 {
		return models.ErrNotFound
	}

	return err
}

func (r *UsersRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, username, password string) (*models.User, error) {
	var user models.User
	if err := r.col.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}
