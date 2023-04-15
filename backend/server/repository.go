package server

import (
	"database/sql"
)

type Repo struct {
	DB *sql.DB
}

func CreateTable(repo *Repo) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
		password TEXT NOT NULL
		);`
	_, err := repo.DB.Exec(query)
	return err
}

func (r *Repo) CreateUser(user *User) error {
	// Get a new dynamo client
	query := `INSERT INTO users (email, name, password)
	VALUES(?, ?, ?);`
	if _, err := r.DB.Exec(query, user.Email, user.Name, user.Password); err != nil {
		return err
	}
	return nil
}

func (r *Repo) Login(login *Login) (User, error) {
	query := `SELECT * FROM users
	WHERE ? = email AND ? = password`
	var user User
	err := r.DB.QueryRow(query, login.Email, login.Password).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	return user, err
}
