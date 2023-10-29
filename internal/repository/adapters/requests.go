package adapters

import (
	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
)

func UserLoginRequestFromAPI(req *api.UserLoginRequest) *models.UserLoginRequest {
	out := &models.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	return out
}

func CreateUserRequestFromAPI(req *api.CreateUserRequest) *models.CreateUserRequest {
	out := &models.CreateUserRequest{
		Username: req.Username,
	}

	return out
}
