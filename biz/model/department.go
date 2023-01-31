package model

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID         uint64         `gorm:"column:id"`
	Name       string         `gorm:"column:name"`
	HospitalID uint64         `gorm:"column:hospital_id"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Department) TableName() string {
	return "departments"
}
