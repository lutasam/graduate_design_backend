package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     uint64             `json:"user_id" bson:"user_id"`
	UserName   string             `json:"user_name" bson:"user_name"`
	UserAvatar string             `json:"user_avatar" bson:"user_avatar"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

func (Message) DBName() string {
	return "messages"
}
