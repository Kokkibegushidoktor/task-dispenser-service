package adapters

import (
	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
)

func UserToApi(user *models.User) *api.User {
	out := &api.User{
		Username: user.Username,
		Password: user.Password,
		Admin:    user.Admin,
	}
	return out
}
