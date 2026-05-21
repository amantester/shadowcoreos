package services

import (
	"github.com/amantester/shadowcoreos/apps/core-api/internal/auth"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/models"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/repositories"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(email, username, password string) (*models.User, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		Username:     username,
		PasswordHash: hashedPassword,
		Role:         "researcher",
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	valid := auth.CheckPassword(password, user.PasswordHash)
	if !valid {
		return "", err
	}

	return auth.GenerateToken(user.ID, user.Role)
}
