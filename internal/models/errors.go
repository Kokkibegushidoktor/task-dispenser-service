package models

import "errors"

var (
	ErrNotFound      = errors.New("error not found")
	ErrAlreadyExists = errors.New("error already exists")
)
