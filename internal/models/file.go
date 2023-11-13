package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	FileStatus int
	FileType   string
)

const (
	ClientUploadInProgress FileStatus = iota
	UploadedByClient
	ClientUploadError
	StorageUploadInProgress
	UploadedToStorage
	StorageUploadError
)

const (
	Image FileType = "image"
	Video FileType = "video"
	Other FileType = "other"
)

type File struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Type        FileType           `json:"type" bson:"type"`
	ContentType string             `json:"contentType" bson:"contentType"`
	Size        int64              `json:"size" bson:"size"`
	Status      FileStatus         `json:"status" bson:"status,omitempty"`
	URL         string             `json:"url" bson:"url,omitempty"`
	Uploader    string             `json:"uploader" bson:"uploader"`
}
