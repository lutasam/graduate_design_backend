package handler

import (
	"github.com/gin-gonic/gin"
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
	}
}

func (ins *HospitalController) FindHospitals(c *gin.Context) {
	req := &bo.FindHospitalsRequest{}
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
