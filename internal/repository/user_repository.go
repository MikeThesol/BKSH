package repository

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByID(id int64) (*models.User, error) {
	return nil, nil
}
