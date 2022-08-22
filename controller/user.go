package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := &models.ParamSignUp{} // 是不是等价于 new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid params ", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			Respi
		}
	}

	// 2. 业务处理
	logic.SignUp()

	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}
