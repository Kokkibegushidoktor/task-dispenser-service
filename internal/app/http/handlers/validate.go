package handlers

import (
	"errors"
	"strings"
)

func validateSignInInput(inp signInInput) error {
	if strings.TrimSpace(inp.Username) == "" &&
		strings.TrimSpace(inp.Password) == "" {
		return errors.New("empty request")
	}

	if strings.TrimSpace(inp.Username) == "" {
		return errors.New("empty username")
	}

	if strings.TrimSpace(inp.Password) == "" {
		return errors.New("empty password")
	}

	return nil
}
