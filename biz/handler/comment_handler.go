package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type CommentController struct{}

func RegisterCommentController(r *gin.RouterGroup) {
	commentController := &CommentController{}
	{
		r.POST("add_comment", commentController.AddComment)
		r.POST("find_comments", commentController.FindComments)
	}
}

func (ins *CommentController) AddComment(c *gin.Context) {
	req := &bo.AddCommentRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetCommentService().AddComment(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *CommentController) FindComments(c *gin.Context) {
	req := &bo.FindCommentsRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetCommentService().FindComments(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
