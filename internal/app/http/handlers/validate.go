package handlers

import (
	"github.com/go-playground/validator/v10"
)

type InputValidator struct {
	validator *validator.Validate
}

func (v *InputValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewInputValidator() *InputValidator {
	return &InputValidator{
		validator: validator.New(),
	}
}
