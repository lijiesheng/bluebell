package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"fmt"
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

func Login(p *models.ParamLogin) (user *models.User,err error) {
	// 1. 判断用户存不存在
	user = &models.User{
		Username : p.Username,
		Password : p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	var token string
	if token, err = jwt.GenToken(user.UserID, user.Username, user.Password); err != nil {
		fmt.Printf("生成 token 失败, err = %s\n", err)
		return nil, err
	}
	user.Token = token    //todo 没有将 token 保存到数据库中
	return user, nil
}