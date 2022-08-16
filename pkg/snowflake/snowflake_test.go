package snowflake

import (
	"fmt"
	"testing"
)

// 一个便于理解的代码 https://blog.csdn.net/bestzy6/article/details/125598962

func TestInit(t *testing.T) {
	if err := Init("2022-08-16", 1); err != nil {
		fmt.Println("init failed , ", err)
		return
	}
	id := GenID()
	fmt.Println(id)
}
