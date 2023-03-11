package services

import (
	"context"
	"github.com/Ekod/highload-otus/domain/users"
)

type Services struct {
	UserService UserService
}

func New(userService UserService) *Services {
	return &Services{
		UserService: userService,
	}
}

type UserService interface {
	LoginUser(ctx context.Context, user *users.UserRequest) (*users.UserResponse, error)
	RegisterUser(ctx context.Context, user *users.UserRequest) (*users.UserResponse, error)
	GetUsers(ctx context.Context) ([]users.UserResponse, error)
	GetCurrentUser(ctx context.Context, userId int) (*users.UserResponse, error)
	GetUsersByFullName(ctx context.Context, firstName, lastName string) ([]users.UserResponse, error)
}
