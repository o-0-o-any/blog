package main

import (
	"blog/db"
	"blog/routers"
)

func main() {
	// 数据库初始化
	db.InitDB()

	// 启动Web Serve  强制绑定localhost与127.0.0.1
	routers.InitRouters().Run("0.0.0.0:8080")
}
