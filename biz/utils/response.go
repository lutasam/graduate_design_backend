package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
)

// Response DIY your Response based on bo.BaseResponse
func Response(c *gin.Context, code, errCode int, errMsg string, data interface{}) {
	resp := &bo.BaseResponse{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	}
	c.JSON(code, resp)
}

// ResponseSuccess returns a 200 code.
// If you need to return other code, use Response func.
func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, common.STATUSOKCODE, common.STATUSOKCODE, common.STATUSOKMSG, data)
}

// ResponseClientError returns a 400 code.
// If you need to return other code, use Response func.
func ResponseClientError(c *gin.Context, err common.Error) {
	Response(c, common.CLIENTERRORCODE, err.Code(), err.Error(), nil)
}

// ResponseServerError returns a 500 code
// If you need to return other code, use Response func.
func ResponseServerError(c *gin.Context, err common.Error) {
	Response(c, common.SERVERERRORCODE, err.Code(), err.Error(), nil)
}

// ResponseError returns an Error
// This function will automatically determine what error type it is, so just use it.
func ResponseError(c *gin.Context, err error) {
	if IsClientError(err) {
		ResponseClientError(c, err.(common.Error))
		return
	} else {
		ResponseServerError(c, err.(common.Error))
		return
	}
}
