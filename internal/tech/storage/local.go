package storage

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path"
)

type LocalFileStorage struct {
	folder string
}

func NewLocalFileStorage(folder string) *LocalFileStorage {
	if err := os.MkdirAll(folder, 0700); err != nil {
		log.Error().Msgf("error setting up local file storage, err: %v", err)
	}
	return &LocalFileStorage{
		folder: folder,
	}
}

func (fs *LocalFileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", fs.folder, path.Dir(input.Name)), 0700); err != nil {
		log.Error().Msgf("error setting up folders for a file, err: %v", err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", fs.folder, input.Name))
	if err != nil {
		return "", err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Error().Msgf("error closing file from Upload(), err: %v", err)
		}
	}()

	_, err = f.ReadFrom(input.File)
	if err != nil {
		return "", err
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *LocalFileStorage) Delete(ctx context.Context, input DeleteInput) error {
	if err := os.Remove(input.Name); err != nil {
		return err
	}

	return nil
}

func (fs *LocalFileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("LOCAL/%s/%s", fs.folder, filename)
}
