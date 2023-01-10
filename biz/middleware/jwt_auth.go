package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/GIN_LUTA/biz/common"
	"github.com/lutasam/GIN_LUTA/biz/utils"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.ResponseClientError(c, common.USERNOTLOGIN)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			token = strings.Split(token, " ")[1]
		}
		jwtStruct, err := utils.ParseJWTToken(token)
		if err != nil {
			utils.ResponseClientError(c, common.EXCEEDTIMELIMIT)
			c.Abort()
			return
		}
		c.Set("jwtStruct", jwtStruct)
		c.Next()
	}
}
