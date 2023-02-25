package bo

import "github.com/lutasam/doctors/biz/vo"

type TakeUserInfoRequest struct {
	UserID string `json:"user_id" binding:"-"`
}

type TakeUserInfoResponse struct {
	User *vo.UserVO `json:"user"`
}

type UpdateUserInfoRequest struct {
	UserID      string `json:"user_id" binding:"-"`
	PhoneNumber string `json:"phone_number" binding:"-"`
	Name        string `json:"name" binding:"-"`
	Avatar      string `json:"avatar" binding:"-"`
	Birthday    string `json:"birthday" binding:"-"`
	Sex         int    `json:"sex" binding:"-"`
	City        string `json:"city" binding:"-"`
	Address     string `json:"address" binding:"-"`
}

type UpdateUserInfoResponse struct{}

type FindUsersRequest struct {
	CurrentPage int    `json:"current_page" binding:"required"`
	PageSize    int    `json:"page_size" binding:"required"`
	Name        string `json:"name" binding:"-"`
}

type FindUsersResponse struct {
	Total int          `json:"total"`
	Users []*vo.UserVO `json:"users"`
}

type DeleteUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type DeleteUserResponse struct{}
