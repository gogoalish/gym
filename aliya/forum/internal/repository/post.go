package repository

import (
	"database/sql"

	"project/internal/model"
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) model.PostRepository {
	return &postRepo{
		db: db,
	}
}
