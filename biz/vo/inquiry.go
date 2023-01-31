package vo

type InquiryTitleVO struct {
	InquiryID               string `json:"inquiry_id"`
	DiseaseName             string `json:"disease_name"`
	ReplyDoctorHospitalName string `json:"reply_doctor_hospital_name"`
	ReplyDoctorName         string `json:"reply_doctor_name"`
}

type InquiryVO struct {
	InquiryID               string `json:"inquiry_id"`
	UserName                string `json:"user_name"`
	DiseaseName             string `json:"disease_name"`
	Description             string `json:"description"`
	WeightHeight            string `json:"weight_height"`
	HistoryOfAllergy        string `json:"history_of_allergy"`
	PastMedicalHistory      string `json:"past_medical_history"`
	OtherInfo               string `json:"other_info"`
	ReplyDoctorName         string `json:"reply_doctor_name"`
	ReplyDoctorHospitalName string `json:"reply_doctor_hospital_name"`
	ReplySuggestion         string `json:"reply_suggestion"`
	CreatedAt               string `json:"created_at"`
}
