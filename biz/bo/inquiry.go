package bo

import "github.com/lutasam/doctors/biz/vo"

type CreateInquiryRequest struct {
	DiseaseName        string `json:"disease_name" binding:"required"`  // 疾病
	Description        string `json:"description" binding:"required"`   // 疾病描述
	WeightHeight       string `json:"weight_height" binding:"-"`        // 身高体重
	HistoryOfAllergy   string `json:"history_of_allergy" binding:"-"`   // 过敏史
	PastMedicalHistory string `json:"past_medical_history" binding:"-"` // 既往病史
	OtherInfo          string `json:"other_info" binding:"-"`           // 其他信息
}

type CreateInquiryResponse struct{}

type DeleteInquiryRequest struct {
	InquiryID string `json:"inquiry_id" binding:"required"`
}

type DeleteInquiryResponse struct{}

type UploadReplySuggestionRequest struct {
	InquiryID       string `json:"inquiry_id" binding:"required"`
	ReplySuggestion string `json:"reply_suggestion" binding:"required"`
}

type UploadReplySuggestionResponse struct{}

type FindInquiryTitlesRequest struct {
	CurrentPage int    `json:"current_page" binding:"required"`
	PageSize    int    `json:"page_size" binding:"required"`
	DiseaseName string `json:"disease_name" binding:"-"`
	ReplyStatus int    `json:"reply_status" binding:"required"` // 1: 已回复 2: 未回复 3: 全部状态
}

type FindInquiryTitlesResponse struct {
	Total         int                  `json:"total"`
	InquiryTitles []*vo.InquiryTitleVO `json:"inquiry_titles"`
}

type FindInquiryRequest struct {
	InquiryID string `json:"inquiry_id" binding:"required"`
}

type FindInquiryResponse struct {
	Inquiry *vo.InquiryVO `json:"inquiry"`
}
