package storage

import (
	"context"
	"io"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

type DeleteInput struct {
	Name string
}

type Provider interface {
	Upload(ctx context.Context, input UploadInput) (string, error)
	Delete(ctx context.Context, input DeleteInput) error
}
