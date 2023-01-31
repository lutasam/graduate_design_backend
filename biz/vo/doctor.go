package vo

type DoctorVO struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	HospitalName     string `json:"hospital_name"`
	DepartmentName   string `json:"department_name"`
	ProfessionalRank string `json:"professional_rank"`
	StudyDirection   string `json:"study_direction"`
	Description      string `json:"description"`
}
