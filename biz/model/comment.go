package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName  string             `json:"user_name" bson:"user_name"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func (Comment) DBName() string {
	return "comments"
}
