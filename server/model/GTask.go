package model

import "time"

type GTask struct {
	ID     uint
	UserID uint
	GID    uint
	Start  time.Time
	Finish time.Time
	Plan   string
}
