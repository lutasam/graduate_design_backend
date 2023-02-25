package vo

type HospitalVO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Address      string `json:"address"`
	HospitalRank string `json:"hospital_rank"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
}
