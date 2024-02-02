package repository

import (
	"context"
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LevelsRepo struct {
	col *mongo.Collection
}

func NewLevelsRepo(db *mongo.Database) *LevelsRepo {
	return &LevelsRepo{
		col: db.Collection("levels"),
	}
}

func (r *LevelsRepo) Create(ctx context.Context, level *models.TaskLevel) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, level)

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *LevelsRepo) Update(ctx context.Context, inp UpdateLevelInput) error {
	updateQuery := bson.M{}

	if inp.Title != "" {
		updateQuery["title"] = inp.Title
	}

	if inp.VarQuestCount != 0 {
		updateQuery["varQuestCount"] = inp.VarQuestCount
	}

	res, err := r.col.UpdateByID(ctx, inp.ID, updateQuery)
	if res.MatchedCount == 0 {
		return models.ErrNotFound
	}

	return err
}

func (r *LevelsRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	res, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount == 0 {
		return models.ErrNotFound
	}

	return err
}

func (r *LevelsRepo) DeleteByTaskId(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.DeleteMany(ctx, bson.M{"taskId": id})

	return err
}

func (r *LevelsRepo) GetByTaskId(ctx context.Context, id primitive.ObjectID) ([]models.TaskLevel, error) {
	var levels []models.TaskLevel

	cur, err := r.col.Find(ctx, bson.M{"taskId": id})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &levels)

	return levels, err
}

func (r *LevelsRepo) GetById(ctx context.Context, id primitive.ObjectID) (*models.TaskLevel, error) {
	var level models.TaskLevel
	if err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&level); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNotFound
		}

		return nil, err
	}
	return &level, nil
}

func (r *LevelsRepo) AddQuestion(ctx context.Context, id primitive.ObjectID, question *models.LevelQuestion) error {
	res, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$push": bson.M{"questions": question}})
	if res.MatchedCount == 0 {
		return models.ErrNotFound
	}

	return err
}

func (r *LevelsRepo) UpdateQuestion(ctx context.Context, inp *models.LevelQuestion) error {
	updateQuery := bson.M{}

	if inp.Title != "" {
		updateQuery["questions.$.title"] = inp.Title
	}

	if inp.Description != "" {
		updateQuery["questions.$.description"] = inp.Description
	}

	if inp.ContentURL != "" {
		updateQuery["questions.$.contentURL"] = inp.ContentURL
	}

	res, err := r.col.UpdateOne(ctx, bson.M{"questions._id": inp.ID}, bson.M{"$set": updateQuery})
	if res.MatchedCount == 0 {
		return models.ErrNotFound
	}

	return err
}

func (r *LevelsRepo) DeleteQuestion(ctx context.Context, id primitive.ObjectID) error {
	res, err := r.col.UpdateOne(ctx, bson.M{"questions._id": id}, bson.M{"$pull": bson.M{"questions": bson.M{"_id": id}}})
	if res.MatchedCount == 0 {
		return models.ErrNotFound
	}
	return err
}
