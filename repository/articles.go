package repository

import (
	"blog/db"
	"blog/models"
	"fmt"
	"strconv"
)

func AddArticles(article models.ArticlesModel) error {
	// 向表中插入博客数据
	err := db.DB.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllArticles() ([]models.ArticlesModel, error) {
	articles := make([]models.ArticlesModel, 0, 10)

	err := db.DB.Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func OrderByIDSearchElem(id string) (models.ArticlesModel, error) {
	bid, result := stringTURNint(id)
	if !result {
		return models.ArticlesModel{}, fmt.Errorf("没有这篇博客")
	}
	article := models.ArticlesModel{}
	err := db.DB.Where("id = ?", bid).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

// 将纯数字字符串转换成int类型
func stringTURNint(id string) (int, bool) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return -1, false
	}
	return value, true
}
