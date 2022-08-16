package mysql

import (
	"fmt"
	"testing"
)

// 测试连接
func TestInit(t *testing.T) {
	err := initMysql()
	if err != nil {
		fmt.Println("mysql 数据连接不上")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("mysql ping 有问题")
	}
	fmt.Println("数据库已经连接上了")
}

//
