package bo

import "github.com/lutasam/doctors/biz/vo"

type AddCommentRequest struct {
	CommentAreaName string `json:"comment_area_name" binding:"required"`
	Content         string `json:"content" binding:"required"`
}

type AddCommentResponse struct {
	Comment *vo.CommentVO `json:"comment"`
}

type FindCommentsRequest struct {
	CommentAreaName string `json:"comment_area_name" binding:"required"`
}

type FindCommentsResponse struct {
	Total    int             `json:"total"`
	Comments []*vo.CommentVO `json:"comments"`
}
