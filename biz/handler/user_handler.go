package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type UserController struct{}

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := &UserController{}
	{
		r.POST("/update_user_info", userController.UpdateUserInfo)
		r.POST("/take_user_info", userController.TakeUserInfo)
		r.POST("/find_users", userController.FindUsers)
		r.POST("/delete_user", userController.DeleteUser)
	}
}

func (ins *UserController) TakeUserInfo(c *gin.Context) {
	req := &bo.TakeUserInfoRequest{}
	resp, err := service.GetUserService().TakeUserInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) UpdateUserInfo(c *gin.Context) {
	req := &bo.UpdateUserInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().UpdateUserInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) FindUsers(c *gin.Context) {
	req := &bo.FindUsersRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().FindUsers(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) DeleteUser(c *gin.Context) {
	req := &bo.DeleteUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().DeleteUser(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
