package repository

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"fmt"
)

func CheckUser(id string) bool {
	user := models.UserModel{}
	db.DB.Where("id = ?", id).Find(&user)

	if user.ID == "" {
		return false
	}
	return true
}

func AddUser(user models.UserModel) error {
	// 密码加密
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = db.DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("create user error: %v", err)
	}
	return nil
}

func GetUser(id string) (models.UserModel, error) {
	user := models.UserModel{}
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("get user error: %v", err)
	}
	return user, nil
}

func UpdateUser(user models.UserModel) error {
	// err := db.DB.Save(&user).Error  错误 传入的user没有给主键赋值 会被当成新的记录进行更新
	err := db.DB.Model(&models.UserModel{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserNums() int {
	var count int64
	if err := db.DB.Model(&models.UserModel{}).Count(&count).Error; err != nil {
		return -1
	}
	return int(count)
}

func JudgeNameisExit(name string) bool {
	var count int64
	db.DB.Model(&models.UserModel{}).Where("name = ?", name).Count(&count)
	return count > 0
}
