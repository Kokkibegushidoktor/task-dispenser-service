package service

import (
	"context"
	"fmt"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/pkg/storage"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
)

var folders = map[models.FileType]string{
	models.Image: "images",
	models.Video: "videos",
	models.Other: "other",
}

type FileService struct {
	storage storage.Provider
	repo    repository.Files
}

func NewFileService(storage storage.Provider, repo repository.Files) *FileService {
	return &FileService{
		storage: storage,
		repo:    repo,
	}
}

func (s *FileService) Save(ctx context.Context, file *models.File) (primitive.ObjectID, error) {
	return s.repo.Create(ctx, file)
}

func (s *FileService) SaveAndUpload(ctx context.Context, file *models.File) (string, error) {
	defer deleteFile(file.Name)

	id, err := s.Save(ctx, file)
	if err != nil {
		return "", err
	}

	url, err := s.upload(ctx, file)
	if err != nil {
		if err := s.repo.UpdateStatus(ctx, id, models.StorageUploadError); err != nil {
			return "", err
		}
		return "", err
	}

	if err = s.repo.UpdateStatusAndSetURL(ctx, id, url); err != nil {
		return "", err
	}

	return url, nil
}

func (s *FileService) upload(ctx context.Context, file *models.File) (string, error) {
	f, err := os.Open(file.Name)
	if err != nil {
		return "", err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Error().Msgf("error closing file from upload(), err: %v", err)
		}
	}()

	return s.storage.Upload(ctx, storage.UploadInput{
		File:        f,
		Name:        s.generateFilename(file),
		ContentType: file.ContentType,
		Size:        file.Size,
	})
}

func (s *FileService) generateFilename(file *models.File) string {
	filename := fmt.Sprintf("%s.%s.%s", file.Uploader, uuid.New().String(), getFileExtension(file.Name))

	return fmt.Sprintf("%s/%s", folders[file.Type], filename)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")

	return parts[len(parts)-1]
}

func deleteFile(name string) {
	if err := os.Remove(name); err != nil {
		log.Error().Msgf("Error deleting file, err: %v", err)
	}
}
