package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	// 2、生成 UID
	userId := snowflake.GenID()
	// 构造一个 User 实例
	user := &models.User{
		UserID: userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}
