package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type DoctorController struct{}

func RegisterDoctorRouter(r *gin.RouterGroup) {
	doctorController := &DoctorController{}
	{
		r.POST("/take_doctor_info", doctorController.TakeDoctorInfo)
		r.POST("/update_doctor_info", doctorController.UpdateDoctorInfo)
		r.POST("/find_doctors", doctorController.FindDoctors)
		r.POST("/delete_doctor", doctorController.DeleteDoctor)
		r.POST("/active_doctor", doctorController.ActiveDoctor)
	}
}

func (ins *DoctorController) TakeDoctorInfo(c *gin.Context) {
	req := &bo.TakeDoctorInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().TakeDoctorInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) UpdateDoctorInfo(c *gin.Context) {
	req := &bo.UpdateDoctorInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().UpdateDoctorInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) FindDoctors(c *gin.Context) {
	req := &bo.FindDoctorsRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().FindDoctors(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) DeleteDoctor(c *gin.Context) {
	req := &bo.DeleteDoctorRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().DeleteDoctor(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) ActiveDoctor(c *gin.Context) {
	req := &bo.ActiveDoctorRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().ActiveDoctor(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
