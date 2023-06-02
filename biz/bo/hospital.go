package bo

import "github.com/lutasam/doctors/biz/vo"

type FindHospitalsRequest struct {
	CurrentPage    int    `json:"current_page" binding:"required"`
	PageSize       int    `json:"page_size" binding:"required"`
	HospitalName   string `json:"hospital_name" binding:"-"`
	City           string `json:"city" binding:"-"`
	HospitalRank   int    `json:"hospital_rank" binding:"-"`
	Characteristic string `json:"characteristic" binding:"-"`
}

type FindHospitalsResponse struct {
	Total     int              `json:"total"`
	Hospitals []*vo.HospitalVO `json:"hospitals"`
}

type TakeHospitalInfoRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
}

type TakeHospitalInfoResponse struct {
	Hospital *vo.HospitalVO `json:"hospital"`
}

type FindHospitalDepartmentNamesRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
}

type FindHospitalDepartmentNamesResponse struct {
	Total           int                `json:"total"`
	DepartmentNames []*vo.DepartmentVO `json:"departments"`
}

type CreateHospitalRequest struct {
	Name           string             `json:"name" binding:"required"`
	City           string             `json:"city" binding:"required"`
	Address        string             `json:"address" binding:"required"`
	HospitalRank   int                `json:"hospital_rank" binding:"required"`
	Description    string             `json:"description" binding:"required"`
	Departments    []*vo.DepartmentVO `json:"departments" binding:"required"`
	Characteristic string             `json:"characteristic" binding:"required"`
}

type CreateHospitalResponse struct{}

type UpdateHospitalInfoRequest struct {
	HospitalID     string `json:"hospital_id" binding:"required"`
	Name           string `json:"name" binding:"-"`
	City           string `json:"city" binding:"-"`
	Address        string `json:"address" binding:"-"`
	HospitalRank   int    `json:"hospital_rank" binding:"-"`
	Description    string `json:"description" binding:"-"`
	Characteristic string `json:"characteristic" binding:"-"`
}

type UpdateHospitalInfoResponse struct{}

type DeleteHospitalRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
}

type DeleteHospitalResponse struct{}

type UpdateHospitalRateScoreRequest struct {
	HospitalID string  `json:"hospital_id" binding:"required"`
	Score      float64 `json:"score" binding:"required"`
}

type UpdateHospitalRateScoreResponse struct{}

type TakeHospitalRankRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
	Area       string `json:"area" binding:"-"`
}

type TakeHospitalRankResponse struct {
	Rank int `json:"rank"`
}

type SetHospitalReadCountRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
}

type SetHospitalReadCountResponse struct{}

type TakeHospitalReadCountRequest struct {
	HospitalID string `json:"hospital_id" binding:"required"`
}

type TakeHospitalReadCountResponse struct {
	ReadCount int `json:"read_count"`
}
