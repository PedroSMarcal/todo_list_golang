package coll

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Task struct {
		ID        primitive.ObjectID `bson:"_id" json:"id"`
		Content   string             `bson:"content, omitempty" json:"content"`
		Finish    bool               `bson:"finish, omitempty" json:"finish"`
		CreatedAt time.Time          `bson:"created_at, omitempty" json:"created_at"`
	}
)
