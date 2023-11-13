package repository

import (
	"context"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesRepo struct {
	db *mongo.Collection
}

func NewFilesRepo(db *mongo.Database) *FilesRepo {
	return &FilesRepo{
		db: db.Collection("files"),
	}
}

func (r *FilesRepo) Create(ctx context.Context, file *models.File) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, file)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *FilesRepo) UpdateStatus(ctx context.Context, id primitive.ObjectID, status models.FileStatus) error {
	_, err := r.db.UpdateByID(ctx, id, bson.M{"$set": bson.M{"status": status}})

	return err
}

func (r *FilesRepo) GetForUploading(ctx context.Context) (*models.File, error) {
	var file models.File

	res := r.db.FindOneAndUpdate(ctx, bson.M{"status": models.UploadedByClient}, bson.M{"$set": bson.M{"status": models.StorageUploadInProgress}})
	err := res.Decode(&file)

	return &file, err
}

func (r *FilesRepo) UpdateStatusAndSetURL(ctx context.Context, id primitive.ObjectID, url string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"url": url, "status": models.UploadedToStorage}})

	return err
}

func (r *FilesRepo) GetByID(ctx context.Context, id, schoolId primitive.ObjectID) (*models.File, error) {
	var file models.File

	res := r.db.FindOne(ctx, bson.M{"_id": id, "schoolId": schoolId})
	err := res.Decode(&file)

	return &file, err
}
