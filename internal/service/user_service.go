package service

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/MikeThesol/BKSH/internal/repository"
)

type UserService struct {
	repository *repository.Repository
}

func NewUserService(repository *repository.Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) GetUserByID(id int64) *models.UserResponse {
	return nil
}
