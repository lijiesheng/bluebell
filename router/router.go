package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
}