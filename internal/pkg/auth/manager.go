package auth

import (
	"errors"
	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	NewJWT(user *api.User, ttl time.Duration) (string, error)
	Check(accessToken *jwt.Token) error
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(user *api.User, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["admin"] = user.Admin

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Check(token *jwt.Token) error {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	claim := claims["admin"]
	if claim == nil {
		return errors.New("missing claim")
	}

	admin, ok := claim.(bool)
	if !ok {
		return errors.New("invalid claim")
	}

	if !admin {
		return errors.New("forbidden")
	}

	return nil
}
