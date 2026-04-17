package main

import (
	"blog/db"
	"blog/routers"
	"blog/utils"
)

func main() {
	// 日志初始化
	utils.InitLogger()
	utils.Logger.Info("日志初始化完成")

	// 数据库初始化
	db.InitDB()

	// 启动Web Serve  强制绑定localhost与127.0.0.1
	routers.InitRouters().Run("0.0.0.0:8080")
	utils.Logger.Info("Web Serve 启动完成")
}
