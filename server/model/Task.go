package model

import "time"

type Task struct {
	ID     uint
	UserID uint
	Start  time.Time
	Finish time.Time
	Plan   string
}
