package domain

import (
	"time"

	"gorm.io/gorm"
)

// Information はお知らせ情報を表すエンティティ。
type Information struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string
	Description string
	Priority    int            `gorm:"default:-10"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Information) TableName() string {
	return "information"
}
