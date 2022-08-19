package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 设置默认值
	viper.SetDefault("fileDir", "./")

	// 读取配置文件
	viper.SetConfigFile("./config.yaml")  // 指定配置文件文件
	viper.SetConfigName("config")  // 配置文件名称(无扩展名) config.yaml
	viper.SetConfigType("yaml")    // 如果配置文件的名称中没有扩展名，需要配置此项  config.yaml
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	viper.AddConfigPath(".")              // 工作目录中查找配置
	err := viper.ReadInConfig()           // 处理读配置文件的错误
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}

}
