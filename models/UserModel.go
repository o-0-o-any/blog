package models

import "time"

type UserModel struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"default: '暂未设置'"`
	Password  string
	Role      int `gorm:"default: 1"`
	CreatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}
