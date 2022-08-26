package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"net/http"
)

// 前端传过来的参数做校验
type SignUpParam struct {
	Age uint8 `json:"age" binding:"gte=1,lte=130"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=13,min=1"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 定义一个全局翻译器 T
var trans ut.Translator

func main() {
	router := gin.Default()
	router.POST("/check_param", func(context *gin.Context) {
		var p SignUpParam
		err := context.ShouldBindJSON(&p)  // post 请求窜如一个 json 的 body
		fmt.Printf("p == %+v", p)
		if err != nil {
			fmt.Printf("err:%s\n", err)
			context.JSON(http.StatusOK, gin.H{
				"msg" :"bindjson报错了",
				"data": gin.H{},
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg" :"没有错",
				"data": p,
			})
		}
	})
	router.Run(":8081")
}

func InitTrans(local string) (err error) {

}