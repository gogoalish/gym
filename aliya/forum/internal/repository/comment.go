package repository

import (
	"database/sql"

	"project/internal/model"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) model.CommentRepository {
	return &commentRepo{
		db: db,
	}
}
