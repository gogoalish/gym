package repository

import (
	"database/sql"

	"project/internal/model"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) model.CategoryRepository {
	return &categoryRepo{
		db: db,
	}
}
