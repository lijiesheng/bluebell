package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 框架操作 Cookie
// todo 一次性获取所有的Cookie
func main() {
	router := gin.Default()
	// 添加 Cookie
	router.POST("/addCookie", func(context *gin.Context) {
		var cookie string
		var err error
		// 获取 Cookie 没有 cookie err 会出现错误
		if cookie, err = context.Cookie("ljs"); err != nil {
			// 没有 cookie 设置 cookie
			context.SetCookie("ljs", "love hsm", 3600 * 24 , "/", "localhost", false, true)
			context.JSON(http.StatusOK, gin.H{
				"msg" : fmt.Sprintf("添加 cookie 成功"),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg" : fmt.Sprintf("添加 cookie 失败, 原来的 cookie 是 %s", cookie),
		})
		return
	})
	// 删除 Cookie
	router.DELETE("/deleteCookie", func(context *gin.Context) {
		var cookie string
		var err error
		// 获取 Cookie 没有 cookie err 会出现错误
		if cookie, err = context.Cookie("ljs"); err != nil {
			// 没有 cookie, 不用删除
			context.JSON(http.StatusOK, gin.H{
				"msg" : fmt.Sprintf("没有您需要删除的 cookie "),
			})
			return
		}
		context.SetCookie("ljs", "love hsm", -1 , "/", "localhost", false, true)
		context.JSON(http.StatusOK, gin.H{
			"msg" : fmt.Sprintf("删除 cookie 成功, cookie_name = %s, cookie_value = %s", "ljs", cookie),
		})
		return
	})
	
	// 修改 Cookie
	router.PUT("/updateCookie", func(context *gin.Context) {
		// 获取 Cookie
		var err error
		if _, err = context.Cookie("ljs"); err != nil {
			// 没有 cookie
			context.JSON(http.StatusOK, gin.H{
				"msg" : fmt.Sprintf("找不到您需要修改的 cookie_name = %s", "ljs"),
			})
			return
		}
		context.SetCookie("ljs", "love 个屁", 3600 * 12, "/", "localhost", false, true)
		context.JSON(http.StatusOK, gin.H{
			"msg" : fmt.Sprintf("cookie 修改成功 cookie_name= %s, cookie_value = %s", "ljs", "love 个屁"),
		})
		return
	})
	
	// 获取 Cookie
	router.GET("/getCookie", func(context *gin.Context) {
		var cookie string
		var err error
		if cookie, err = context.Cookie("ljs"); err != nil {
			// 没有获取到 cookie
			context.JSON(http.StatusOK, gin.H{
				"msg" : fmt.Sprintf("没有您获取的cookie_name = %s", "ljs"),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg" : fmt.Sprintf("获取的 cookie 是 %s", cookie),
		})
		return
	})
	router.Run()
}