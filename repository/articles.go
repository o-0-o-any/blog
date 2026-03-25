package repository

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"fmt"
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
	bid, err := utils.StringTOInt(id)
	if err != nil {
		return models.ArticlesModel{}, fmt.Errorf("没有这篇博客")
	}
	article := models.ArticlesModel{}
	err = db.DB.Where("id = ?", bid).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}
