package user_service

import (
	"context"

	"github.com/Ekod/otus-highload/domain"
	"github.com/Ekod/otus-highload/internal/dto"
	"github.com/Ekod/otus-highload/utils/security"
)

type UserRepository interface {
	LoginUser(context.Context, *dto.SecurityUser) (*domain.User, error)
	GetCurrentUser(context.Context, int) (*dto.UserResponse, error)
	GetUsers(context.Context, int) ([]dto.UserResponse, error)
	SaveUser(context.Context, *dto.UserRequest) (int, error)
	GetUsersByFullName(context.Context, string, string) ([]dto.UserResponse, error)
}

type Service struct {
	userRepository      UserRepository
	userRepositorySlave UserRepository
}

func New(userRepository UserRepository, userRepositorySlave UserRepository) *Service {
	return &Service{
		userRepository:      userRepository,
		userRepositorySlave: userRepositorySlave,
	}
}

func (s *Service) LoginUser(ctx context.Context, user *dto.SecurityUser) (*dto.UserResponse, error) {
	foundUser, err := s.userRepository.LoginUser(ctx, user)
	if err != nil {
		return nil, err
	}

	if err = security.VerifyPassword(user.Password, foundUser.Password); err != nil {
		return nil, err
	}

	token, err := security.GenerateToken(foundUser.ID)
	if err != nil {
		return nil, err
	}

	responseUser := dto.UserResponse{
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
		Email:     foundUser.Email,
		Interests: foundUser.Interests,
		City:      foundUser.City,
		Age:       foundUser.Age,
		Gender:    foundUser.Gender,
		Token:     token,
	}

	return &responseUser, nil
}

func (s *Service) RegisterUser(ctx context.Context, user *dto.UserRequest) (*dto.UserResponse, error) {
	hp, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hp

	userId, err := s.userRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := security.GenerateToken(userId)
	if err != nil {
		return nil, err
	}

	responseUser := dto.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Interests: user.Interests,
		City:      user.City,
		Age:       user.Age,
		Gender:    user.Gender,
		Token:     token,
	}

	return &responseUser, nil
}

func (s *Service) GetUsers(ctx context.Context, id int) ([]dto.UserResponse, error) {
	usersList, err := s.userRepository.GetUsers(ctx, id)
	if err != nil {
		return nil, err
	}

	return usersList, nil
}

func (s *Service) GetCurrentUser(ctx context.Context, userId int) (*dto.UserResponse, error) {
	user, err := s.userRepositorySlave.GetCurrentUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUsersByFullName(ctx context.Context, firstName, lastName string) ([]dto.UserResponse, error) {
	usersList, err := s.userRepositorySlave.GetUsersByFullName(ctx, firstName, lastName)
	if err != nil {
		return nil, err
	}

	return usersList, nil
}
