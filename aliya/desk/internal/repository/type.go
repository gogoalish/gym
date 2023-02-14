package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type typeRepo struct {
	db *dbBase
}

func NewTypeRepository(db *sql.DB) TypeRepository {
	return &typeRepo{
		db: newDB(db),
	}
}

type TypeRepository interface {
	GetTypeByID(int) (*model.Type, error)
	GetTypeByName(string) (*model.Type, error)
}

func (t *typeRepo) GetTypeByID(id int) (*model.Type, error) {
	task_type := model.Type{}
	query := `SELECT * FROM task_types WHERE id = ?`
	row, err := t.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&task_type.ID, &task_type.Name); err != nil {
		return nil, err
	}

	return &task_type, nil
}

func (t *typeRepo) GetTypeByName(name string) (*model.Type, error) {
	task_type := model.Type{}
	query := `SELECT * FROM task_types WHERE name = ?`
	row, err := t.db.queryRow(query, name)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&task_type.ID, &task_type.Name); err != nil {
		return nil, err
	}

	return &task_type, nil
}
