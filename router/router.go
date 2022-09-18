package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}

	r := gin.New()

	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1") // 得到一个路由组

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	// 上面的路由不受这个中间件控制
	v1.Use(middlewares.JWTAuthMiddleware())

	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "验证 jwt 成功",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// 根据时间或分数获取帖子列表
	v1.GET("/post2", controller.GetPostListHandler2)

	//r.NoRoute(func(context *gin.Context) { // 没有找到路由
	//	fmt.Println("没有找到路由")
	//	context.HTML(http.StatusNotFound, "views/404.html", nil)
	//})
	return r
}
