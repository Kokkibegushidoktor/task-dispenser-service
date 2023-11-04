package repository

import (
	"context"
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TasksRepo struct {
	col *mongo.Collection
}

func NewTasksRepo(db *mongo.Database) *TasksRepo {
	return &TasksRepo{
		col: db.Collection("tasks"),
	}
}

func (r *TasksRepo) Create(ctx context.Context, task *models.Task) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, task)

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *TasksRepo) Update(ctx context.Context, inp UpdateTaskInput) error {
	updateQuery := bson.M{}

	if inp.Title != "" {
		updateQuery["title"] = inp.Title
	}

	if inp.Description != "" {
		updateQuery["description"] = inp.Description
	}

	_, err := r.col.UpdateByID(ctx, inp.ID, bson.M{"$set": updateQuery})

	return err
}

func (r *TasksRepo) GetById(ctx context.Context, taskId primitive.ObjectID) (*models.Task, error) {
	var task *models.Task
	if err := r.col.FindOne(ctx, bson.M{"_id": taskId}).Decode(task); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNotFound
		}

		return nil, err
	}
	return task, nil
}

func (r *TasksRepo) Delete(ctx context.Context, taskId primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": taskId})

	return err
}
