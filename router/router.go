package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	r := gin.New()

	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))


	v1 := r.Group("/api/v1")   // 得到一个路由组

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/login", controller.LoginHandler)
	// 上面的路由不受这个中间件控制
	v1.Use(middlewares.JWTAuthMiddleware())



	// 根据时间或分数获取帖子列表
	v1.GET("/post2", controller.GetPostListHandler2)



















	r.NoRoute(func(context *gin.Context) {   // 没有找到路由
		fmt.Println("没有找到路由")
		context.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	return r
}