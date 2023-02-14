package repository

import (
	"database/sql"

	"project/internal/model"
)

type reactionRepo struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) model.ReactionRepository {
	return &reactionRepo{
		db: db,
	}
}
