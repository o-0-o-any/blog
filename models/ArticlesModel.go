package models

import "time"

type ArticlesModel struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Title       string
	RedirectURL string //`gorm:"-"` // 不在表中创建这个字段 暂时取消
	Author      string
	Date        time.Time
	Text        string `gorm:"type:longtext"`
}

func (ArticlesModel) TableName() string {
	return "articles"
}
