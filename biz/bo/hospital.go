package bo

import "github.com/lutasam/doctors/biz/vo"

type FindHospitalsRequest struct {
	CurrentPage  int    `json:"current_page" binding:"required"`
	PageSize     int    `json:"page_size" binding:"required"`
	HospitalName string `json:"hospital_name" binding:"-"`
	City         string `json:"city" binding:"-"`
	HospitalRank int    `json:"hospital_rank" binding:"-"`
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
	Total           int      `json:"total"`
	DepartmentNames []string `json:"department_names"`
}
