package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.tmpl", nil)
}

func PerMessageHandler(c *gin.Context) {
	// 获取访问的博客id
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
