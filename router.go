package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/handler"
	"github.com/lutasam/doctors/biz/middleware"
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
	// 登录模块
	login := r.Group("/login")
	handler.RegisterLoginRouter(login)

	// 通讯模块
	talk := r.Group("/talk", middleware.JWTAuth())
	handler.RegisterTalkRouter(talk)

	// 用户模块
	user := r.Group("/user", middleware.JWTAuth())
	handler.RegisterUserRouter(user)

	// 医生模块
	doctor := r.Group("/doctor", middleware.JWTAuth())
	handler.RegisterDoctorRouter(doctor)

	// 问诊模块
	inquiry := r.Group("/inquiry", middleware.JWTAuth())
	handler.RegisterInquiryRouter(inquiry)
}
