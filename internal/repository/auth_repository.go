package repository

import (
	"github.com/MikeThesol/BKSH/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRespository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRespository {
	return &AuthRespository{
		db: db,
	}
}

func (a *AuthRespository) CreateUser(user *models.User) (int64, error) {
	return 0, nil
}
