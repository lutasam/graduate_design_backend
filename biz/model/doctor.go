package model

import (
	"gorm.io/gorm"
	"time"
)

type Doctor struct {
	ID               uint64         `gorm:"column:id"`
	UserID           uint64         `gorm:"column:user_id"`
	User             *User          `gorm:"foreignKey:user_id;references:id"`
	HospitalID       uint64         `gorm:"hospital_id"`
	Hospital         *Hospital      `gorm:"foreignKey:hospital_id;references:id"`
	DepartmentID     uint64         `gorm:"column:department_id"`
	Department       *Department    `gorm:"foreignKey:department_id;references:id"`
	ProfessionalRank int            `gorm:"column:professional_rank"`
	StudyDirection   string         `gorm:"column:study_direction"`
	Description      string         `gorm:"column:description"`
	IsActive         bool           `gorm:"column:is_active"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Doctor) TableName() string {
	return "doctors"
}
