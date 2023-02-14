package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type sprintRepo struct {
	db *dbBase
}

func NewSprintRepository(db *sql.DB) SprintRepository {
	return &sprintRepo{
		db: newDB(db),
	}
}

type SprintRepository interface {
	CreateSprint(*model.Sprint) error
	GetOneSprint(int) (*model.Sprint, error)
	GetAllSprintsOfProject(int) ([]model.Sprint, error)
	UpdateSprint(*model.Sprint) error
	DeleteSprint(int) error
}

//TODO: sprint creation date
func (s *sprintRepo) CreateSprint(sprint *model.Sprint) error {
	query := `INSERT INTO sprints (name, created, expired, description, project_id) VALUES(?,?,?,?,?)`
	return s.db.execOne(query, sprint.Name, sprint.Created, sprint.Expired, sprint.Description, sprint.Project.ID)
}

//TODO
func (s *sprintRepo) GetOneSprint(id int) (*model.Sprint, error) {
	sprint := model.Sprint{}
	query := `SELECT * FROM sprints WHERE id = ?`
	row, err := s.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}
	if err := row.Scan(&sprint.ID, &sprint.Name, &sprint.Created, &sprint.Expired, &sprint.Description, &sprint.Project.ID ); err != nil {
		return nil, err
	}

	return &sprint, nil
}

func (s *sprintRepo) GetAllSprintsOfProject(project_id int) ([]model.Sprint, error) {
	query := `SELECT * FROM sprints WHERE project_id = ?`
	rows, err := s.db.query(query, project_id)
	if err != nil {
		return nil ,err
	}
	defer rows.Close()
	var sprints []model.Sprint

	for rows.Next() {
		var sprint model.Sprint
		if err := rows.Scan(&sprint.ID, &sprint.Name, &sprint.Created, &sprint.Expired, &sprint.Description, &sprint.Project.ID); err != nil {
			return nil, err
		}
		sprints = append(sprints, sprint)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sprints, nil

}
//* Can be updated (name, expired, description)
func (s *sprintRepo) UpdateSprint(sprint *model.Sprint) error {
	query := `UPDATE sprints SET name = ?, expired = ?, description = ?`
	return s.db.execOne(query, sprint.Name, sprint.Expired, sprint.Description)
}

func (s *sprintRepo) DeleteSprint(id int) error {
	query := `DELETE FROM sprints WHERE id = ?`
	return s.db.execOne(query, id)
}
