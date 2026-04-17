package db

import (
	"blog/config"
	"blog/models"
	"blog/utils"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Error error
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Configs.Database.User,
		config.Configs.Database.Password,
		config.Configs.Database.Host,
		config.Configs.Database.Port,
		config.Configs.Database.DBName)
	DB, Error = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 判断连接成功
	if Error != nil {
		// 后续待加日志操作...
		fmt.Println(Error)
		panic(Error)
	}

	// 根据项目中所需的Model自动在数据库中建表 实现表的自动迁移
	if err := DB.AutoMigrate(&models.UserModel{}, &models.ArticlesModel{}); err != nil {
		utils.Logger.Error("数据库表迁移失败", zap.Error(err))
		panic(err)
	}
	utils.Logger.Info("数据库表迁移成功")
}
