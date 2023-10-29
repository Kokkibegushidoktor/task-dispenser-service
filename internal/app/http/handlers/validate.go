package handlers

import (
	"errors"
	"strings"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
)

func validateUserLoginRequest(req *models.UserLoginRequest) error {
	if strings.TrimSpace(req.Username) == "" &&
		strings.TrimSpace(req.Password) == "" {
		return errors.New("empty request")
	}

	if strings.TrimSpace(req.Username) == "" {
		return errors.New("empty username")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("empty password")
	}

	return nil
}
