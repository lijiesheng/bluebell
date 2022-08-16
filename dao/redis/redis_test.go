package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	err := initRedis()
	if err != nil {
		fmt.Println("redis 连接错误", err)
	} else {
		fmt.Println("redis 连接成功")
	}
}
