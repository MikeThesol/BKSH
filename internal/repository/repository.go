package repository

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user *models.User) (int64, error)
}

type User interface {
	GetUserByID(id int64) (*models.User, error)
}

type Book interface {
	CreateBook(book models.Book, userID int64) (int64, error)
}

type Repository struct {
	Auth
	User
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
		User: NewUserRepository(db),
		Book: NewBookRepository(db),
	}
}
