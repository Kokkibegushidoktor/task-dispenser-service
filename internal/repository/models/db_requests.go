package models

type UserLoginRequest struct {
	Username string
	Password string
}

type CreateUserRequest struct {
	Username string
}
