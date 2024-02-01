package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type FileStorage struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewFileStorage(client *minio.Client, bucket, endpoint string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := fs.client.PutObject(ctx, fs.bucket, input.Name, input.File, input.Size, opts)
	if err != nil {
		return "", err
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *FileStorage) Delete(ctx context.Context, input DeleteInput) error {
	return errors.New("NOT IMPLEMENTED")
}

func (fs *FileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("https://%s/%s/%s", fs.endpoint, fs.bucket, filename)
}
