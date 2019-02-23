package model

import (
	"time"
)

type Information struct {
	ID          uint
	Name        string
	Description string
	Priority    int `gorm:"default:-10"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
