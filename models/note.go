package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:size:255`
	Content   string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
