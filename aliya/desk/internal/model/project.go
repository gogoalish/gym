package model

import "time"

type Project struct {
	ID          int
	Title       string
	Description string
	Created     time.Time
	Srints      []Sprint
}


