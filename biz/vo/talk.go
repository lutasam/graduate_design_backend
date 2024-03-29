package vo

type TalkedUserVO struct {
	ID            string `json:"id"`
	Avatar        string `json:"avatar"`
	Name          string `json:"name"`
	LastMessage   string `json:"last_message"`
	PhoneNumber   string `json:"phone_number"`
	CreatedAt     string `json:"created_at"`
	MessageNumber int    `json:"message_number"`
	IsOnline      bool   `json:"is_online"`
}

type MessageVO struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	ImageURL   string `json:"image_url"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
}
