package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/GIN_LUTA/biz/bo"
	"github.com/lutasam/GIN_LUTA/biz/common"
	"github.com/lutasam/GIN_LUTA/biz/service"
	"github.com/lutasam/GIN_LUTA/biz/utils"
)

type DemoController struct{}

func RegisterDemoRouter(r *gin.RouterGroup) {
	demoController := &DemoController{}
	{
		r.GET("/ping", demoController.Ping)
		r.POST("/hello", demoController.Hello)
	}
}

func (ins *DemoController) Ping(c *gin.Context) {
	pong, err := service.GetDemoService().Ping(c)
	if err != nil {
		utils.ResponseServerError(c, common.UNKNOWNERROR)
		return
	}
	utils.ResponseSuccess(c, pong)
}

func (ins *DemoController) Hello(c *gin.Context) {
	req := &bo.HelloRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	hello, err := service.GetDemoService().Hello(c, req)
	if err != nil {
		utils.ResponseServerError(c, common.UNKNOWNERROR)
		return
	}
	utils.ResponseSuccess(c, hello)
}
