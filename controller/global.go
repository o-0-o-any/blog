package controller

import (
	"blog/models"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handler(name string) func(c *gin.Context, data any) {
	return func(c *gin.Context, data any) {
		// 解析组合模板
		t, err := template.ParseFiles("templates/index/layout.tmpl", name)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		// 渲染组合模板
		err = t.ExecuteTemplate(c.Writer, "layout", data)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	}
}

func IndexHandler(c *gin.Context) {
	f := handler("templates/index/index.tmpl")
	f(c, nil)
}

func MessageHandler(c *gin.Context) {
	emp := []models.ArticlesModel{
		{
			Id:          "1",
			Title:       "测试",
			RedirectURL: "/text",
			Author:      "s",
			Date:        time.Date(2026, 3, 20, 10, 30, 0, 0, time.Local),
			Text:        "Hello World! 测试内容文章!",
		},
	}
	handler("templates/index/message.tmpl")(c, gin.H{
		"articles": emp,
	})
}

func DriverHandler(c *gin.Context) {
	handler("templates/index/driver.tmpl")(c, nil)
}
