package coll

import (
	"time"

	"gorm.io/gorm"
)

type (
	Task struct {
		ID        string         `bson:"id, omitempty" json:"id"`
		Name      string         `bson:"name, omitempty" json:"name"`
		Content   string         `bson:"content, omitempty" json:"content"`
		EndDate   time.Time      `bson:"enddate, omitempty" json:"enddate"`
		CreatedAt time.Time      `bson:"created_at, omitempty" json:"created_at"`
		DeletedAt gorm.DeletedAt `bson:"deleted_at, omitempty" json:"deleted_at"`
	}
)
