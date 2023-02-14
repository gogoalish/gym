package model

import "time"

type Sprint struct {
	ID          int
	Project     *Project
	Name        string
	Created     time.Time
	Expired     time.Time
	Description string
	Tasks       []Task
}
