package models

import "time"

type UserModel struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"default: '暂未设置'" json:"name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      int       `gorm:"default: 1" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserModel) TableName() string {
	return "users"
}
