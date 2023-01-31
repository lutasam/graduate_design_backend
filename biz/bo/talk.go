package bo

import "github.com/lutasam/doctors/biz/vo"

type GetTalkedUsersRequest struct{}

type GetTalkedUsersResponse struct {
	Total int              `json:"total"`
	Users []*vo.TalkedUser `json:"users"`
}

type SendMessageRequest struct {
	MessageType int    `json:"message_type"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserAvatar  string `json:"user_avatar"`
	Content     string `json:"content"`
}
