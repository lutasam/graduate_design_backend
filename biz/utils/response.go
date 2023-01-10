package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/GIN_LUTA/biz/bo"
	"github.com/lutasam/GIN_LUTA/biz/common"
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

// ResponseError returns a ServerErrorResponse.
// User ResponseServerError func and ResponseClientError func to distinguish between them.
func ResponseError(c *gin.Context, err common.Error) {
	ResponseServerError(c, err)
}
