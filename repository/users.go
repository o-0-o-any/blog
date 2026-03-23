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
