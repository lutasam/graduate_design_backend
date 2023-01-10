package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/GIN_LUTA/biz/handler"
	"io"
	"os"
)

func InitRouterAndMiddleware(r *gin.Engine) {
	// 设置log文件输出
	logFile, err := os.Create("log/log.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 注册全局中间件Recovery和Logger
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 注册分组路由
	// /demo
	demo := r.Group("/demo")
	handler.RegisterDemoRouter(demo)
}
