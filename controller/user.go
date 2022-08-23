package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := &models.ParamSignUp{} // 是不是等价于 new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid params ", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		_, ok := err.(validator.ValidationErrors)    // 验证参数是否传入错误
		if !ok {
			ResponseError(c, CodeInvalidParam)   // 请求参数错误
			return
		}
		ResponseError(c, CodeInvalidParam)
		//ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))   // todo 这个暂时不知道啥意思
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		ResponseErrorWithError(c, err)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
