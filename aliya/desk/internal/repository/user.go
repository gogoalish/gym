package repository

import (
	"database/sql"

	"service-desk/internal/model"
)

type userRepo struct {
	db *dbBase
}

type UserRepository interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	GetUser(id int) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	DeleteUser(id int) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{
		db: newDB(db),
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	query := `INSERT INTO users (first_name, last_name, description, email, password, location, phone) VALUES(?,?,?,?,?,?,?)`
	return u.db.execOne(query, user.First_name, user.Last_name, user.Description, user.Email, user.Password, user.Location, user.Phone)
}

func (u *userRepo) UpdateUser(user *model.User) error {
	query := `
	UPDATE users
	SET first_name = ?, last_name = ?, description = ?, email = ?, password = ?, location = ?, phone = ?, 
	WHERE id = ?`

	id := user.ID
	return u.db.execOne(query, user.First_name, user.Last_name, user.Description, user.Email, user.Password, user.Location, user.Phone, id)
}

func (u *userRepo) GetUser(id int) (*model.User, error) {
	user := model.User{}

	query := `SELECT * FROM users WHERE id = ?`

	row, err := u.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := row.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Description, &user.Email, &user.Location, &user.Phone); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) DeleteUser(id int) error {
	query := `
	DELETE FROM users 
	WHERE id = ?`

	return u.db.execOne(query, id)
}

func (u *userRepo) GetAllUsers() ([]model.User, error) {
	rows, err := u.db.query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var users []model.User

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Description, &user.Email, &user.Location, &user.Phone); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
