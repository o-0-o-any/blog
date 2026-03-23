package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}

var Configs = &Config{}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// 将配置文件信息解析到结构体中
	if err := viper.Unmarshal(Configs); err != nil {
		fmt.Println(err)
	}
}
