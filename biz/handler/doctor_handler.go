package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		r.POST("/find_hospital_doctors", doctorController.FindHospitalDoctors)
		r.POST("/update_doctor_rate_score", doctorController.UpdateDoctorRateScore)
		r.POST("/take_doctor_rank", doctorController.TakeDoctorRank)
		r.POST("/set_doctor_read_count", doctorController.SetDoctorReadCount)
		r.POST("/take_doctor_read_count", doctorController.TakeDoctorReadCount)
	}
}

func (ins *DoctorController) TakeDoctorInfo(c *gin.Context) {
	req := &bo.TakeDoctorInfoRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
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
	err := c.ShouldBindBodyWith(req, binding.JSON)
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
	err := c.ShouldBindBodyWith(req, binding.JSON)
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
	err := c.ShouldBindBodyWith(req, binding.JSON)
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
	err := c.ShouldBindBodyWith(req, binding.JSON)
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

func (ins *DoctorController) FindHospitalDoctors(c *gin.Context) {
	req := &bo.FindHospitalDoctorsRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().FindHospitalDoctors(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) UpdateDoctorRateScore(c *gin.Context) {
	req := &bo.UpdateDoctorRateScoreRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().UpdateDoctorRateScore(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) TakeDoctorRank(c *gin.Context) {
	req := &bo.TakeDoctorRankRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().TakeDoctorRank(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) SetDoctorReadCount(c *gin.Context) {
	req := &bo.SetDoctorReadCountRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().SetDoctorReadCount(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DoctorController) TakeDoctorReadCount(c *gin.Context) {
	req := &bo.TakeDoctorReadCountRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDoctorService().TakeDoctorReadCount(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
