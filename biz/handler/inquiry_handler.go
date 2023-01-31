package handler

import (
	"github.com/gin-gonic/gin"
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
	}
}

func (ins *InquiryController) CreateInquiry(c *gin.Context) {
	req := &bo.CreateInquiryRequest{}
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
	err := c.ShouldBind(req)
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
