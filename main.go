package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/setting"
	"fmt"
)

func main() {
	// 加载配置 setting
	if err := setting.Init("./conf/config.yaml"); err != nil {
		fmt.Printf("load config failed, err: %s\n", err)
		return
	}
	// 加载日志配置 logger.Init
	if err := logger.Init(setting.Conf.Log, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 加载 mysql mysql.Init   mysql 要 close
	if err := mysql.InitMysql(setting.Conf.Mysql); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()  // todo

	// 加载 redis redis.Init   redis 要 close
	if err := redis.InitRedis(setting.Conf.Redis); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
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
