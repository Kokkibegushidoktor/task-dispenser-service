package models

import "errors"

var (
	ErrNotFound         = errors.New("error not found")
	ErrAlreadyExists    = errors.New("error already exists")
	ErrPassAlreadySetUp = errors.New("error password already setup")
)
