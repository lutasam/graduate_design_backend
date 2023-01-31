package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/middleware"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type LoginController struct{}

func RegisterLoginRouter(r *gin.RouterGroup) {
	loginController := &LoginController{}
	{
		r.POST("/login", loginController.Login)
		r.POST("/apply_register", loginController.ApplyRegister)
		r.POST("/active_user", loginController.ActiveUser)
		r.POST("/reset_password", loginController.ResetPassword)
		r.POST("/active_reset_password", loginController.ActiveResetPassword)
		r.POST("/apply_change_user_email", middleware.JWTAuth(), loginController.ApplyChangeUserEmail)
		r.POST("/active_change_user_email", middleware.JWTAuth(), loginController.ActiveChangeUserEmail)
	}
}

func (ins *LoginController) Login(c *gin.Context) {
	req := &bo.LoginRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().Login(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ApplyRegister(c *gin.Context) {
	req := &bo.ApplyRegisterRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().ApplyRegister(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ActiveUser(c *gin.Context) {
	req := &bo.ActiveUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveUser(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ResetPassword(c *gin.Context) {
	req := &bo.ResetPasswordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().ResetPassword(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ActiveResetPassword(c *gin.Context) {
	req := &bo.ActiveResetPasswordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveResetPassword(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ApplyChangeUserEmail(c *gin.Context) {
	req := &bo.ApplyChangeUserEmailRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ApplyChangeUserEmail(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ActiveChangeUserEmail(c *gin.Context) {
	req := &bo.ActiveChangeUserEmailRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveChangeUserEmail(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
