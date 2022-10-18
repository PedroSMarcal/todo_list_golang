package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Task struct {
		ID          uint `gorm:"primaryKey; autoIncrement; unique; <-:false"`
		Name        string
		Description string
		Author      string
		LastDay     time.Time
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}
)
