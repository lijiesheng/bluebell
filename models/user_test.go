package models

import (
	"bluebell/dao/mysql"
	"fmt"
	"testing"
	"time"
)

func init()  {
	mysql.InitMysql()
}

func TestBatchInsertUsers(t *testing.T) {
	// 注意传入的参数是 []interface{}
	user := make([]User, 1)
	user[0] = User{
		User_id:     100,
		Username:    "名称10",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	user = append(user, User{
		User_id:     200,
		Username:    "名称200",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})
	user = append(user, User{
		User_id:     300,
		Username:    "名称300",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})
	user2 := []interface{}{user[0],user[1], user[2]}          // 结构体体数组 ==> 转换为 interface[] 数组
	err := BatchInsertUsers(user2)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err : %v\n", err)
	}
}

func TestBatchInsertUser3(t *testing.T) {
	user := make([]User, 1)
	user[0] = User{
		User_id:     101,
		Username:    "名称101",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	user = append(user, User{
		User_id:     210,
		Username:    "名称210",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})
	user = append(user, User{
		User_id:     301,
		Username:    "名称310",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})
	u := []*User{&user[0], &user[1], &user[2]}
	err := BatchInsertUser3(u)
	if err != nil {
		fmt.Printf("sqlx batch insert failed , err : %v", err)
	}
}

func TestQueryByIDs(t *testing.T) {
	ids := []int{2, 3, 4,9,12,13}
	user, err := QueryByIDs(ids)
	if err != nil {
		fmt.Printf("queryByIDs failed err : %v", err)
	}
	for i:=0; i < len(user); i++ {
		fmt.Printf("%+v\n", user[i])
	}
}

