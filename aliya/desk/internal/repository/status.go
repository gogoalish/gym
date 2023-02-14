package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type statusRepo struct {
	db *dbBase
}

func NewStatusRepository(db *sql.DB) StatusRepository {
	return &statusRepo{
		db: newDB(db),
	}
}

type StatusRepository interface {
	GetStatusByID(int) (*model.Status, error)
	GetStatusByName(string) (*model.Status, error)
}

func (s *statusRepo) GetStatusByID(id int) (*model.Status, error) {
	status := model.Status{}
	query := `SELECT * FROM task_statuses WHERE id = ?`
	row, err := s.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&status.ID, &status.Name); err != nil {
		return nil, err
	}

	return &status, nil
}

// TODO : make status name unique
func (s *statusRepo) GetStatusByName(name string) (*model.Status, error) {
	status := model.Status{}
	query := `SELECT * FROM task_statuses WHERE name = ?`
	row, err := s.db.queryRow(query, name)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&status.ID, &status.Name); err != nil {
		return nil, err
	}

	return &status, nil
}
