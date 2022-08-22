package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context)  {
	// 1. 获取参数和参数校验
	c.Query("")


	// 2. 业务处理
	logic.SignUp()

	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}