package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	r := gin.New()

	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	// 登录
	r.GET("/login", controller.LoginHandler)


	r.NoRoute(func(context *gin.Context) {   // 没有找到路由
		fmt.Println("没有找到路由")
		context.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	return r
}