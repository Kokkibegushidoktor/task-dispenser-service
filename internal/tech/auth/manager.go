package auth

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	NewJWT(user *models.User, ttl time.Duration) (string, error)
	Check(accessToken *jwt.Token) error
	GetRoles(token *jwt.Token) ([]string, error)
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

func (m *Manager) Check(token *jwt.Token) error {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	claim := claims["adm"]
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

func (m *Manager) GetRoles(token *jwt.Token) ([]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}

	claim := claims["rls"]
	if claim == nil {
		return nil, errors.New("empty claim")
	}

	rls, ok := claim.([]interface{})
	if !ok {
		return nil, errors.New("invalid claim")
	}

	roles := make([]string, len(rls))

	for i, role := range rls {
		roles[i], ok = role.(string)
		if !ok {
			return nil, errors.New("invalid claim")
		}
	}

	return roles, nil
}
