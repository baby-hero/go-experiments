package services

import (
	"context"
	"userapp/models"
	"userapp/repositories"
)

type UserService interface {
	RegisterUser(ctx context.Context, input models.RegisterInput) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	UpdateUser(ctx context.Context, id uint, input models.UpdateInput) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, input models.RegisterInput) (*models.User, error) {
	// You can add business rules here (e.g., check if email already exists)
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password, // Hash it before saving in production
	}
	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, input models.UpdateInput) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	return s.userRepo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}
