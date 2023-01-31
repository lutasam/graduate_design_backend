package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/utils"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	InitRouterAndMiddleware(r)

	err := r.Run(":" + utils.GetConfigString("server.port"))
	if err != nil {
		panic(err)
	}
}
