package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type priorityRepo struct {
	db *dbBase
}

func NewPriorityRepository(db *sql.DB) PriorityRepository {
	return &priorityRepo{
		db: newDB(db),
	}
}

type PriorityRepository interface {
	GetPriorityByID(int) (*model.Priority, error)
	GetPriorityByName(string) (*model.Priority, error)
}

func (p *priorityRepo) GetPriorityByID(id int) (*model.Priority, error) {
	priority := model.Priority{}
	query := `SELECT * FROM priority WHERE id = ?`
	row, err := p.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&priority.ID, &priority.Name); err != nil {
		return nil, err
	}

	return &priority, nil
}

func (t *priorityRepo) GetPriorityByName(name string) (*model.Priority, error) {
	priority := model.Priority{}
	query := `SELECT * FROM priority WHERE name = ?`
	row, err := t.db.queryRow(query, name)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&priority.ID, &priority.Name); err != nil {
		return nil, err
	}

	return &priority, nil
}
