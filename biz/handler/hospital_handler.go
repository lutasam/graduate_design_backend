package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type HospitalController struct{}

func RegisterHospitalRouter(r *gin.RouterGroup) {
	hospitalController := &HospitalController{}
	{
		r.POST("find_hospitals", hospitalController.FindHospitals)
		r.POST("take_hospital_info", hospitalController.TakeHospitalInfo)
		r.POST("find_hospital_department_names", hospitalController.FindHospitalDepartmentNames)
		r.POST("create_hospital", hospitalController.CreateHospital)
		r.POST("update_hospital_info", hospitalController.UpdateHospitalInfo)
		r.POST("delete_hospital", hospitalController.DeleteHospital)
		r.POST("update_hospital_rate_score", hospitalController.UpdateHospitalRateScore)
		r.POST("/take_hospital_rank", hospitalController.TakeHospitalRank)
		r.POST("/set_hospital_read_count", hospitalController.SetHospitalReadCount)
		r.POST("/take_hospital_read_count", hospitalController.TakeHospitalReadCount)
	}
}

func (ins *HospitalController) FindHospitals(c *gin.Context) {
	req := &bo.FindHospitalsRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().FindHospitals(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) TakeHospitalInfo(c *gin.Context) {
	req := &bo.TakeHospitalInfoRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().TakeHospitalInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) FindHospitalDepartmentNames(c *gin.Context) {
	req := &bo.FindHospitalDepartmentNamesRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().FindHospitalDepartmentNames(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) CreateHospital(c *gin.Context) {
	req := &bo.CreateHospitalRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().CreateHospital(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) UpdateHospitalInfo(c *gin.Context) {
	req := &bo.UpdateHospitalInfoRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().UpdateHospitalInfo(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) DeleteHospital(c *gin.Context) {
	req := &bo.DeleteHospitalRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().DeleteHospital(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) UpdateHospitalRateScore(c *gin.Context) {
	req := &bo.UpdateHospitalRateScoreRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().UpdateHospitalRateScore(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) TakeHospitalRank(c *gin.Context) {
	req := &bo.TakeHospitalRankRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().TakeHospitalRank(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) SetHospitalReadCount(c *gin.Context) {
	req := &bo.SetHospitalReadCountRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().SetHospitalReadCount(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *HospitalController) TakeHospitalReadCount(c *gin.Context) {
	req := &bo.TakeHospitalReadCountRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetHospitalService().TakeHospitalReadCount(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
