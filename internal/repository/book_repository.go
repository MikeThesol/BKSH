package repository

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) CreateBook(book models.Book, userID int64) (int64, error) {
	return 0, nil
}
