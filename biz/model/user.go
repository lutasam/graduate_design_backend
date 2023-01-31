package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint64         `gorm:"column:id"`
	Email         string         `gorm:"column:email"`
	PhoneNumber   string         `gorm:"column:phone_number"`
	Password      string         `gorm:"column:password"`
	Name          string         `gorm:"column:name"`
	Birthday      time.Time      `gorm:"column:birthday"`
	Avatar        string         `gorm:"column:avatar"`
	CharacterType int            `gorm:"column:character_type"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
