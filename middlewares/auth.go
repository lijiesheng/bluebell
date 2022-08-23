package middlewares

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于 JWT 的认证中间件


// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "" {
			// 1、放在请求头中
			// 这里假设Token放在Header的Authorization中，并使用Bearer开头
			// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
			// 这里的具体实现方式要依据你的实际业务情况决定

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				controller.ResponseError(c, controller.CodeInvalidToken)
				c.Abort()
				return
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
			token, err := jwt.ParseToken(parts[1])
			if err != nil {
				controller.ResponseError(c, controller.CodeInvalidToken)
				c.Abort()
				return
			}
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set(controller.CtxUserIDKey, token.UserID)
			// 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
			c.Next()
		}
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
	}

	// 2、放在请求体中


	// 3、放在 URI 中

}











