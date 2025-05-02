package service

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/MikeThesol/BKSH/internal/repository"
)

type Auth interface {
	RegisterUser(user *models.User) (int64, error)
}

type User interface {
	GetUserByID(id int64) *models.UserResponse
}

type Book interface {
	GetBookByID(id int64) *models.Book
}

type Service struct {
	Auth
	User
	Book
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository),
		User: NewUserService(repository),
		Book: NewBookService(repository),
	}
}
