package service

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/MikeThesol/BKSH/internal/repository"
)

type BookService struct {
	repository *repository.Repository
}

func NewBookService(repository *repository.Repository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (b *BookService) GetBookByID(id int64) *models.Book {
	return nil
}
