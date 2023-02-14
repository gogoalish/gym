package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type projectRepo struct {
	db *dbBase
}

type ProjectRepository interface {
	GetAllProjectsOfUser(int) ([]model.Project, error)
	GetOneProject(int) (*model.Project, error)
	CreateProject(*model.Project) error
	UpdateProject(*model.Project) error
	DeleteProject(int) error
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepo{
		db: newDB(db),
	}
}

// TODO: creation date
func (p *projectRepo) CreateProject(project *model.Project) error {
	query := `INSERT INTO projects (title, description, created) VALUES(?, ?, ?)`
	return p.db.execOne(query, project.Title, project.Description, project.Created)
}

func (p *projectRepo) GetAllProjectsOfUser(user_id int) ([]model.Project, error) {
	query := `SELECT * FROM projects WHERE project_id = ?`
	rows, err := p.db.query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []model.Project

	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Created); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *projectRepo) GetOneProject(id int) (*model.Project, error) {
	project := model.Project{}
	query := `SELECT * FROM projects WHERE id = ?`
	row, err := p.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&project.ID, &project.Title, &project.Description, &project.Created); err != nil {
		return nil, err
	}

	return &project, nil
}

// * Can be updated only title, description columns
func (p *projectRepo) UpdateProject(project *model.Project) error {
	query := `UPDATE projects SET title = ?, description = ?`
	return p.db.execOne(query, project.Title, project.Description)
}

func (p *projectRepo) DeleteProject(id int) error {
	query := `DELETE FROM projects WHERE id = ?`
	return p.db.execOne(query, id)
}
