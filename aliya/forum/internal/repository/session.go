package repository

import (
	"database/sql"

	"project/internal/model"
)

type sessionRepo struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) model.SessionRepository {
	return &sessionRepo{
		db: db,
	}
}
