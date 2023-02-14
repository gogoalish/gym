package repository

import (
	"database/sql"

	"project/internal/model"
)

type userRepo struct {
	db *dbBase
}

func NewUserRepository(db *sql.DB) model.UserRepository {
	return &userRepo{
		db: newDB(db),
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	query := `
	INSERT INTO Users (name, email, password) 
	VALUES (?,?,?)`

	return u.db.execOne(query, user.Name, user.Email, user.Password)
}

func (u *userRepo) ReadUserById(id string) (*model.User, error) {
	user := model.User{}

	query := `SELECT * FROM Users WHERE id = ?`

	row, err := u.db.queryRow(query, id)
	if err != nil {
		return nil, err
	}

	if err := user.ScanRow(row); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) ReadUserForAuthorization(user model.User) (*model.User, error) {
	query := `SELECT * FROM Users WHERE email = ? AND password = ?`

	row, err := u.db.queryRow(query, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	if err := user.ScanRow(row); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(user model.User, id string) error {
	query := `
	UPDATE Users
	SET name = ?, email = ?, password = ?
	WHERE id = ?`

	return u.db.execOne(query, user.Name, user.Email, user.Password, id)
}

func (u *userRepo) DeleteUser(user model.User, id string) error {
	query := `
	DELETE FROM Users 
	WHERE name= ? AND email=? AND password=?`

	return u.db.execOne(query, user.Name, user.Email, user.Password)
}
