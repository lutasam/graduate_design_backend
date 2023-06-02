package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/service"
	"github.com/lutasam/doctors/biz/utils"
)

type InquiryController struct{}

func RegisterInquiryRouter(r *gin.RouterGroup) {
	inquiryController := &InquiryController{}
	{
		r.POST("/create_inquiry", inquiryController.CreateInquiry)
		r.POST("/delete_inquiry", inquiryController.DeleteInquiry)
		r.POST("/upload_reply_suggestion", inquiryController.UploadReplySuggestion)
		r.POST("/find_inquiry_titles", inquiryController.FindInquiryTitles)
		r.POST("/find_inquiry", inquiryController.FindInquiry)
		r.POST("/find_doctor_inquiries", inquiryController.FindDoctorInquiries)
		r.POST("/find_user_inquiries", inquiryController.FindUserInquiries)
		r.POST("/find_doctor_suggestion_inquiries", inquiryController.FindDoctorSuggestionInquiries)
	}
}

func (ins *InquiryController) CreateInquiry(c *gin.Context) {
	req := &bo.CreateInquiryRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().CreateInquiry(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) DeleteInquiry(c *gin.Context) {
	req := &bo.DeleteInquiryRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().DeleteInquiry(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) UploadReplySuggestion(c *gin.Context) {
	req := &bo.UploadReplySuggestionRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().UploadReplySuggestion(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) FindInquiryTitles(c *gin.Context) {
	req := &bo.FindInquiryTitlesRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().FindInquiryTitles(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) FindInquiry(c *gin.Context) {
	req := &bo.FindInquiryRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().FindInquiry(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) FindDoctorInquiries(c *gin.Context) {
	req := &bo.FindDoctorInquiriesRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().FindDoctorInquiries(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) FindUserInquiries(c *gin.Context) {
	req := &bo.FindUserInquiriesRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().FindUserInquiries(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *InquiryController) FindDoctorSuggestionInquiries(c *gin.Context) {
	req := &bo.FindDoctorSuggestionInquiriesRequest{}
	err := c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetInquiryService().FindDoctorSuggestionInquiries(c, req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}
	utils.ResponseSuccess(c, resp)
}
