package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "@@@@ljs_hsm_com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := DB.Get(&count ,sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	// 执行 SQL 入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = DB.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// encryptPassword 对密码进行加密
func encryptPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}


// Login 登录
func Login(user *models.User) (err error) {
	oldPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	if err := DB.Get(user, sqlStr, user.Username); err != nil {
		if err == sql.ErrNoRows {
			return ErrorUserNotExist
		}
		return err
	}
	// 判断密码是否正确
	if encryptPassword(oldPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}