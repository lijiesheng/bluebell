package controller

import (
	"bluebell/dao/redis"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// 返回参数
type ResUser struct {
	*models.User
	AccessToken string
}

// 注册
func SignUpHandler(c *gin.Context) {
	fmt.Println("进来了")
	// 1. 获取参数和参数校验
	p := &models.ParamSignUp{} // 是不是等价于 new(models.ParamSignUp)
	// post 的 json 格式
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid params ", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		_, ok := err.(validator.ValidationErrors) // 验证参数是否传入错误
		if !ok {
			ResponseError(c, CodeInvalidParam) // 请求参数错误
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

// 登录
func LoginHandler(c *gin.Context) {
	fmt.Println("LoginHandler")
	// 1. 获取参数和参数校验
	p := &models.ParamLogin{}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("LoginHandler with invalid params ", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		_, ok := err.(validator.ValidationErrors) // 验证参数是否传入错误
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseError(c, CodeInvalidParam)
		//ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))   // todo 这个暂时不知道啥意思
		return
	}

	// 2. 业务处理
	var user *models.User
	var err error
	if user, err = logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ResponseErrorWithError(c, err)
		return
	}
	fmt.Printf("%+v\n", user)

	// 2.1 生成 token
	token, err := jwt.GenToken(user.UserID, p.Username, p.Password)
	if err != nil {
		zap.L().Error("生成 token 错误")
		return
	}
	// 将 redis 和 userid 绑定，限制同一账号同一时间只能登录一个机器
	isExistToken := redis.RDB.Get(context.TODO(), strconv.FormatInt(user.UserID, 10)).Val()
	if isExistToken != "" {
		ResponseErrorWithError(c, errors.New("您的设备已经登录, 请退出后然后再此机器上登录"))
		return
	}

	redis.RDB.Set(context.TODO(), strconv.FormatInt(user.UserID, 10), token, 0)

	// 3. 返回响应
	resUser := ResUser{
		User:        user,
		AccessToken: token,
	}
	ResponseSuccess(c, &resUser)
}
