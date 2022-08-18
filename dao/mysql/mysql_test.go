package mysql

import (
	"bluebell/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

func init() {
	err := InitMysql()
	if err != nil {
		fmt.Println("mysql 数据连接不上")
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("mysql ping 有问题")
	}
	fmt.Println("数据库已经连接上了")
}

// 使用 sqlx 结构体的字段名要大写
// 测试查询单条
func TestQueryRowDemo(t *testing.T) {
	sql := "select id, user_id, username, password, email, gender, create_time, update_time from user where id = ?"
	var u models.User
	err := DB.Get(&u, sql, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("%+v", u)
}

// 测试查询多行
func TestQueryMultiRowDemo(t *testing.T) {
	sql := "select id, user_id, username, password, email, gender, create_time, update_time  from user where id > ?"
	var u_list []models.User
	err := DB.Select(&u_list, sql, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", u_list)
}

// 插入
func TestInsertDemo(t *testing.T) {
	// 测试一次性插入一条数据
	sqlstr := "insert into user(user_id, username, password, email, gender, create_time, update_time) " +
		" values(?, ?, ?, ?, ? ,?, ?)"
	now := time.Now()
	ret, err := DB.Exec(sqlstr, 14, "小李子", "", "ljs_hsm@163.com", 0, now, now)
	if err != nil {
		fmt.Printf("insert failed, err : %v\n", err)
		return
	}
	lastInsertId, err := ret.LastInsertId() // 新插入数据的 id
	if err != nil {
		fmt.Printf("get lastInsertId failed, err : %v\n", err)
		return
	}
	fmt.Println("新插入的数据 id : ", lastInsertId)

	//// 测试一次性插入两条数据
	//sqlstr := "insert into user(user_id, username, password, email, gender, create_time, update_time) " +
	//	" values(?, ?, ?, ?, ? ,?, ?), (?, ?, ?, ?, ? ,?, ?)"
	//now := time.Now()
	//params := make([]interface{}, 0)
	//params = append(params, 30, "小李子_1", "", "ljs_hsm@1631.com", 0, now, now, 31, "小李子_2", "", "ljs_hsm@1635.com", 0, now, now)
	//fmt.Println(params...)
	//DB.Exec(sqlstr, params...)
}

// 测试更新
func TestUpdateDemo(t *testing.T) {
	sqlstr := "update user set username = ? where id = ? "
	result, err := DB.Exec(sqlstr, "小狗狗", 5)
	if err != nil {
		fmt.Printf("update failed, err : %v\n", err)
		return
	}
	lastUpdateId, err := result.RowsAffected() // 影响操作行数
	if err != nil {
		fmt.Printf("get lastUpdateId failed, err : %v\n", err)
		return
	}
	fmt.Println("修改了 : ", lastUpdateId, " 行数据")
}

// 测试删除
func TestDeleteDemo(t *testing.T) {
	sqlstr := "delete from user where id = ?"
	result, err := DB.Exec(sqlstr, 5)
	if err != nil {
		fmt.Printf("delete failed, err : %v\n", err)
		return
	}
	deleteAffectId, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get delete failed, err : %v\n", err)
		return
	}
	fmt.Println("删除了 : ", deleteAffectId, " 行数据")
}

// 测试 NamedQuery
// 可以传入结构体 和 map
func TestNamedQuery(t *testing.T) {
	sqlStr := "select * from user where username = :username"
	// 使用 map 做命名查询
	mapstr := map[string]interface{}{
		"username": "李结胜",
	}
	rows, err := DB.NamedQuery(sqlStr, mapstr)
	if err != nil {
		fmt.Printf("DB.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed , err: %v\n", err)
			continue
		}
		fmt.Printf("%+v\n", u)
	}

	// 使用结构体命令查询，根据结构体字段的 db tag 进行映射
	u := models.User{
		Username: "李结胜",
	}
	rows, err = DB.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("DB.NameQuery failed, err : %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed , err: %v\n", err)
			continue
		}
		fmt.Printf("%+v\n", u)
	}
}

// 测试 NamedExec [增删改查]
func TestNamedExec(t *testing.T) {
	sqlstr := "insert into user(user_id, username, password, email, gender, create_time, update_time) " +
		" values(:user_id, :username, :password, :email, :gender , :create_time, :update_time)"
	// 使用 map
	//m := map[string]interface{}{
	//	"user_id":13,
	//	"username":"小黄",
	//	"password":"",
	//	"email": "15827200980@163.com",
	//	"gender": 0,
	//	"create_time": time.Now(),
	//	"update_time": time.Now(),
	//}
	//result, err := DB.NamedExec(sqlstr, m)
	//if err != nil {
	//	fmt.Printf("DB.NamedExec failed, err : %v\n", err)
	//	return
	//}
	//lastInsertId, err := result.LastInsertId()
	//if err != nil {
	//	fmt.Printf("get lastInsertId failed, err : %v\n", err)
	//	return
	//}
	//fmt.Printf("插入数据的 id 是 %d\n", lastInsertId)
	// 使用 map 结束

	// 使用 struct
	u := models.User{
		User_id:     15,
		Username:    "小黄小李",
		Password:    "",
		Email:       "15827509197@163.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	result, err := DB.NamedExec(sqlstr, u)
	if err != nil {
		fmt.Printf("DB.NamedExec failed, err : %v\n", err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId failed, err : %v\n", err)
		return
	}
	fmt.Printf("插入数据的 id 是 %d\n", lastInsertId)
	// 使用 struct 结束
}

// 测试 事务  DB.Beginx(), tx.Exec() tx.Commit() tx.Rollback()
func TestTransactionDemo(t *testing.T) {
	var err error
	tx, err := DB.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err : %v\n", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() // 回滚
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "update user set username = ? where id = ?"
	result, err := tx.Exec(sqlStr1, "我是事务", 1)
	if err != nil {
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		return
	}

	sqlStr2 := "update user set username = ? where i = ?" // 故意把 sql 语句写错，看看是否可以回滚
	result, err = tx.Exec(sqlStr2, "我是事务", 1)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		return
	}
}

// 测试 sqlx.In

// 测试 批量插入 方法一： 自行构造批量插入的语句
//   这个方法费内存
func TestBatchInsertUsersMethod_1(t *testing.T) {
	// 比如我要插入 3 个 user
	userStrings := make([]models.User, 1)
	userStrings[0] = models.User{
		User_id:     20,
		Username:    "名称1",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	userStrings = append(userStrings, models.User{
		User_id:     30,
		Username:    "名称2",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})
	userStrings = append(userStrings, models.User{
		User_id:     40,
		Username:    "名称3",
		Password:    "",
		Email:       "2398027035@qq.com",
		Gender:      0,
		Create_time: time.Now(),
		Update_time: time.Now(),
	})

	//valueArgs := make([]interface{}, 0, 3*7) // 3 个 user 对象，每个 user 对象有 7 个元素, id 不用拼接在 sql 中
	//// 准备数据 这里可以参考

	//for _, v := range userStrings {
	//	valueArgs = append(valueArgs, v.User_id, v.Username, v.Password, v.Email, v.Gender, v.Create_time, v.Update_time)
	//}
	sqlstr := "insert into user(user_id, username, password, email, gender, create_time, update_time) values "
	for i := 0; i < 3-1; i++ {
		sqlstr += "(?, ?, ?, ?, ? ,?, ?), "
	}
	sqlstr += "(?, ?, ?, ?, ? ,?, ?)"
	fmt.Println(sqlstr)
	params := make([]interface{}, 0, 3*7)
	for _, v := range userStrings {
		params = append(params, v.User_id, v.Username, v.Password, v.Email, v.Gender, v.Create_time, v.Update_time)
	}
	DB.Exec(sqlstr, params...)

	//fmt.Printf("%v\n", valueArgs)
	//stms := fmt.Sprintf("insert into user(user_id, username, password, email, gender, create_time, update_time) VALUES ",
	//	strings.Join(valueArgs, ","))
	//fmt.Println("stms == ",stms)
	//fmt.Sprintf("insert into user(user_id, username, password, email, gender, create_time, update_time) values %s",
	//	strings.Join(valueArgs, ","))
}


