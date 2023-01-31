package vo

type UserVO struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Name          string `json:"name"`
	Birthday      string `json:"birthday"`
	Avatar        string `json:"avatar"`
	CharacterType string `json:"character_type"`
}
