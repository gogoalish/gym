package model

type scan interface {
	Scan(agr ...any) error
}
type User struct {
	Id       string
	Email    string
	Name     string
	Password string
}

type UserRepository interface {
	CreateUser(*User) error
}

func (u *User) ScanRow(row scan) error {
	return row.Scan(
		u.Id,
		u.Name,
		u.Email,
		u.Password,
	)
}
