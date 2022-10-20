package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type (
	RequestTask struct {
		Content string
	}

	ResponseTask struct {
		ID        primitive.ObjectID `bson:"_id" json:"id"`
		Content   string             `bson:"content, omitempty" json:"content"`
		CreatedAt time.Time          `bson:"created_at, omitempty" json:"created_at"`
		DeletedAt gorm.DeletedAt     `bson:"deleted_at, omitempty" json:"deleted_at"`
	}
)
