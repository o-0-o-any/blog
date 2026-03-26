package controller

import (
	"blog/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 延时重定向函数 被路由函数调用
func redirectHandler(c *gin.Context, content, redirectURL string) {
	c.HTML(http.StatusOK, "endEdit.tmpl", gin.H{
		"redirectURL": redirectURL,
		"content":     content,
	})
}

func PerMessageHandler(c *gin.Context) {
	// 获取访问的博客id
	id := c.Param("blogID")
	// 根据id在数据库中查找指定的博客
	article, err := repository.OrderByIDSearchElem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"content": "Don't have this blog",
		})
	}
	// 渲染显示博客的界面
	// 防止覆盖/blog/message/edit的界面 使用global中的handler函数间接渲染模板
	handler("templates/message/article.tmpl")(c, gin.H{
		"title":  article.Title,
		"author": article.Author,
		"date":   article.Date,
		"text":   article.Text,
	})
}
