package vo

type HospitalVO struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	City           string  `json:"city"`
	Address        string  `json:"address"`
	HospitalRank   string  `json:"hospital_rank"`
	Description    string  `json:"description"`
	RateScore      float64 `json:"rate_score"`
	RatePeopleNum  int     `json:"rate_people_num"`
	Characteristic string  `json:"characteristic"`
	CreatedAt      string  `json:"created_at"`
}
