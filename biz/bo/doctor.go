package bo

import "github.com/lutasam/doctors/biz/vo"

type TakeDoctorInfoRequest struct {
	DoctorID string `json:"doctor_id" binding:"-"`
}

type TakeDoctorInfoResponse struct {
	Doctor *vo.DoctorVO `json:"doctor"`
}

type UpdateDoctorInfoRequest struct {
	DoctorID         string `json:"doctor_id" binding:"-"`
	HospitalName     string `json:"hospital_name" binding:"-"`
	DepartmentName   string `json:"department_name" binding:"-"`
	ProfessionalRank int    `json:"professional_rank" binding:"-"`
	StudyDirection   string `json:"study_direction" binding:"-"`
	Description      string `json:"description" binding:"-"`
}

type UpdateDoctorInfoResponse struct{}

type FindDoctorsRequest struct {
	CurrentPage    int    `json:"current_page" binding:"required"`
	PageSize       int    `json:"page_size" binding:"required"`
	Name           string `json:"name" binding:"-"`
	HospitalName   string `json:"hospital_name" binding:"-"`
	DepartmentName string `json:"department_name" binding:"-"`
}

type FindDoctorsResponse struct {
	Total   int            `json:"total"`
	Doctors []*vo.DoctorVO `json:"doctors"`
}

type DeleteDoctorRequest struct {
	DoctorID string `json:"doctor_id"`
}

type DeleteDoctorResponse struct{}

type ActiveDoctorRequest struct {
	DoctorID string `json:"doctor_id"`
}

type ActiveDoctorResponse struct{}
