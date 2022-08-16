package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/pkg/snowflake"
	"fmt"
)

func main() {
	// 加载配置 setting

	// 加载日志配置 logger.Init

	// 加载 mysql mysql.Init   mysql 要 close
	err := mysql.InitMysql()
	if err != nil {

	}
	//defer mysql.Close()  // todo

	// 加载 redis redis.Init   redis 要 close
	err = redis.InitRedis()
	if err != nil {

	}

	// 加载 雪花算法
	//snowflake.Init("20")
	if err := snowflake.Init("2022-08-16", 1); err != nil {
		fmt.Println("init failed , ", err)
		return
	}
	id := snowflake.GenID()
	fmt.Println(id)
}
