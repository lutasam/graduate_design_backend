package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/handler"
	"github.com/lutasam/doctors/biz/middleware"
	"io"
	"net/http"
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

	// 虚拟静态文件路径 windows开启 linux注释掉下方代码
	r.StaticFS("/imgs", http.Dir("D:\\Code\\program\\Go\\src\\graduate_design_backend\\imgs"))

	// 注册分组路由
	// 登录模块
	login := r.Group("/login", middleware.SensitiveFilter())
	handler.RegisterLoginRouter(login)

	// 通讯模块
	talk := r.Group("/talk")
	handler.RegisterTalkRouter(talk)

	// 用户模块
	user := r.Group("/user", middleware.JWTAuth(), middleware.SensitiveFilter())
	handler.RegisterUserRouter(user)

	// 医院模块
	hospital := r.Group("/hospital", middleware.JWTAuth(), middleware.SensitiveFilter())
	handler.RegisterHospitalRouter(hospital)

	// 医生模块
	doctor := r.Group("/doctor", middleware.JWTAuth(), middleware.SensitiveFilter())
	handler.RegisterDoctorRouter(doctor)

	// 问诊模块
	inquiry := r.Group("/inquiry", middleware.JWTAuth(), middleware.SensitiveFilter())
	handler.RegisterInquiryRouter(inquiry)

	// 文件模块
	file := r.Group("/file")
	handler.RegisterFileRouter(file)

	// 评论模块
	comment := r.Group("/comment", middleware.JWTAuth(), middleware.SensitiveFilter())
	handler.RegisterCommentController(comment)
}
