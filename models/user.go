package models

import (
	"bluebell/dao/mysql"
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	Id int64
	User_id int64
	Username string
	Password string
	Email string
	Gender int8
	Create_time time.Time
	Update_time time.Time
}

// 实现批量插入方式一的前提
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.User_id, u.Username, u.Password, u.Email, u.Gender, u.Create_time, u.Update_time}, nil
}

// 实现批量插入方式一
func BatchInsertUsers(users []interface{}) error {
	// 如果 args 实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	sqlInsert := "insert into user(user_id, username, password, email, gender, create_time, update_time) values"
	for i := 0; i < len(users) -1; i++ {
		sqlInsert += "(?),"
	}
	sqlInsert += "(?)"
	query, args , err := sqlx.In(sqlInsert, users...,)
	if err != nil {
		fmt.Printf("sqlx batch insert users failed, %+v", err)
	}
	fmt.Println(query)
	fmt.Println(args)
	_, err = mysql.DB.Exec(query, args...)
	return err
}

// 实现批量方式二  使用 NamedExec 实现批量插入
// 注意 sql语句最后不能有空格和;
// 我个人认为这个方法可以，传递的是指针数组，开销较小
func BatchInsertUser3(users []*User) error {
	_, err := mysql.DB.NamedExec("insert into user(user_id, username, password, email, gender, create_time, update_time) VALUES (:user_id, :username, :password, :email, :gender, :create_time, :update_time)", users)
	return err
}



// sqlx.In 查询
func QueryByIDs(ids []int)(users []User, err error) {
	// 动态填充 id
	query, args, err := sqlx.In("select id, user_id, username, password, email, gender, create_time, update_time from user where id in (?)", ids)
	if err != nil {
		return
	}
	query = mysql.DB.Rebind(query)
	err = mysql.DB.Select(&users, query, args...)
	return
}

//

