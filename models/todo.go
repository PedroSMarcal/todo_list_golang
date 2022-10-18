package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Task struct {
		ID        uint `gorm:"primaryKey; autoIncrement; unique; <-:false"`
		Name      string
		Content   string
		Author    string
		EndDate   time.Time
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
