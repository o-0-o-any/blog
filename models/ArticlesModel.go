package models

import "time"

type ArticlesModel struct {
	Id          string `gorm:"primary_key"`
	Title       string
	RedirectURL string
	Author      string
	Date        time.Time
	Text        string `gorm:"type:longtext"`
}

func (ArticlesModel) TableName() string {
	return "articles"
}
