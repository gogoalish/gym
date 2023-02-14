package model

type scan interface {
	Scan(agr ...any) error
}

type User struct {
	ID          int
	First_name  string
	Last_name   string
	Description string
	Email       string
	Password    string
	Location    string
	Phone       string
}

func (u *User) ScanRow(row scan) error {
	return row.Scan(
		u.ID,
		u.First_name,
		u.Last_name,
		u.Description,
		u.Email,
		u.Password,
		u.Location,
		u.Phone,
	)
}
