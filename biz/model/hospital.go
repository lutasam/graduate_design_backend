package model

import (
	"gorm.io/gorm"
	"time"
)

type Hospital struct {
	ID           uint64         `gorm:"column:id"`
	Name         string         `gorm:"column:name"`
	City         string         `gorm:"column:city"`
	Address      string         `gorm:"column:address"`
	HospitalRank int            `gorm:"column:hospital_rank"`
	Description  string         `gorm:"column:description"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Hospital) TableName() string {
	return "hospitals"
}
