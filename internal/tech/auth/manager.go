package auth

import (
	"errors"
	"fmt"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	NewJWT(user *models.User, ttl time.Duration) (string, error)
	Parse(accessToken string) (*jwt.Token, error)
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

func (m *Manager) NewJWT(user *models.User, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["sub"] = user.ID.Hex()
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["adm"] = user.Admin
	claims["rls"] = user.Roles

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
