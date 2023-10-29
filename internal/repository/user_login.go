package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/adapters"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository/models"
)

func (r *RepoImpl) UserLogin(ctx context.Context, req *api.UserLoginRequest) (*api.User, error) {

	filter := userLoginFilter(adapters.UserLoginRequestFromAPI(req))

	var user *models.User
	if err := r.UsersCl.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, api.ErrNotFound
		}
		return nil, fmt.Errorf("error finding user, err: %v", err)
	}

	return adapters.UserToApi(user), nil
}
