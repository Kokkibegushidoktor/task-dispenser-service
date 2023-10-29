package repository

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
)

func userLoginFilter(req *models.UserLoginRequest) bson.M {
	return bson.M{"username": req.Username, "password": req.Password}
}
