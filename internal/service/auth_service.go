package service

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/MikeThesol/BKSH/internal/repository"
)

type AuthService struct {
	repository *repository.Repository
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (a *AuthService) RegisterUser(user *models.User) (int64, error) {
	return 0, nil
}
