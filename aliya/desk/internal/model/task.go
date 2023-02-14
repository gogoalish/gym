package model

type Task struct {
	ID          int
	Title       string
	Description string
	Author      *User
	Project     *Project
	Sprint      *Sprint
	Type        *Type
	Status      *Status
	Priority    *Priority
}



