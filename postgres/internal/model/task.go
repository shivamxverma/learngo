package model

import "time"

type Task struct {
	ID 			int				`json:"id"`
	Title 		string			`json:"title"`
	Completed 	bool
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}