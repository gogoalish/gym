package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type taskRepo struct {
	db *dbBase
}

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	GetOneTask(int) (*model.Task, error)
	CreateTask(*model.Task) error
	UpdateTask(*model.Task) error
	DeleteTask(int) error
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepo{
		db: newDB(db),
	}
}

func (t *taskRepo) GetAllTasks() ([]model.Task, error) {
	rows, err := t.db.query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var tasks []model.Task

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Author.ID, &task.Project.ID, &task.Sprint.ID, &task.Type.ID, &task.Status.ID, &task.Priority.ID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepo) GetOneTask(id int) (*model.Task, error) {
	task := model.Task{}
	query := `SELECT * FROM tasks WHERE id = ?`
	row, err := t.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Author.ID, &task.Project.ID, &task.Sprint.ID, &task.Type.ID, &task.Status.ID, &task.Priority.ID); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepo) CreateTask(task *model.Task) error {
	query := `INSERT INTO tasks (title, description, author_id, type_id, project_id, sprint_id, status_id, priority_id) VALUES(?,?,?,?,?,?,?,?)`
	return t.db.execOne(query, task.Title, task.Description, task.Author.ID, task.Type.ID, task.Project.ID, task.Sprint.ID, task.Status.ID, task.Priority.ID)
}

func (t *taskRepo) UpdateTask(task *model.Task) error {
	query := `UPDATE tasks SET title = ?, description = ?, type_id = ?, sprint_id = ?, status_id = ?, priority_id = ? WHERE id = ?`

	return t.db.execOne(query, task.Title, task.Description, task.Type.ID, task.Sprint.ID, task.Status.ID, task.Priority.ID)
}

func (t *taskRepo) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	return t.db.execOne(query, id)
}
