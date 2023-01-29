package repository

import (
	"database/sql"

	"project/internal/model"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) model.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	return nil
}
