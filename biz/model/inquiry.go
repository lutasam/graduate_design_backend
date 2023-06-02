package model

import (
	"gorm.io/gorm"
	"time"
)

type Inquiry struct {
	ID                 uint64         `gorm:"column:id"`
	UserID             uint64         `gorm:"column:user_id"` // 发起人
	User               User           `gorm:"foreignKey:user_id;references:id"`
	ReplyDoctorID      uint64         `gorm:"column:reply_doctor_id"` // 接诊医生
	Doctor             Doctor         `gorm:"foreignKey:reply_doctor_id;references:id"`
	DiseaseName        string         `gorm:"column:disease_name"`         // 疾病
	Description        string         `gorm:"column:description"`          // 疾病描述
	WeightHeight       string         `gorm:"column:weight_height"`        // 身高体重
	HistoryOfAllergy   string         `gorm:"column:history_of_allergy"`   // 过敏史
	PastMedicalHistory string         `gorm:"column:past_medical_history"` // 既往病史
	OtherInfo          string         `gorm:"column:other_info"`           // 其他信息
	ReplySuggestion    string         `gorm:"column:reply_suggestion"`     // 问诊建议
	CreatedAt          time.Time      `gorm:"column:created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Inquiry) TableName() string {
	return "inquiries"
}

type InquiryES struct {
	InquiryID      uint64    `json:"inquiry_id"`
	Title          string    `json:"title"`
	Describe       string    `json:"describe"`
	DescribeVector []float64 `json:"describe_vector"`
}

func (InquiryES) IndexName() string {
	return "inquiries"
}
