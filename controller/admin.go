package controller

import (
	"blog/models"
	"blog/repository"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "主页",
	})
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func SignUpHandler(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	result := repository.CheckUser(id)
	// 储存到结构体中
	user := models.UserModel{ID: id, Password: password}
	// id不重复时 创建user行数据
	if !result {
		err := repository.AddUser(user)
		if err != nil {
			htmlHandler(c, "注册失败!系统异常!", "/blog/admin/login")
		} else {
			htmlHandler(c, "注册成功!已自动为你登录!", "/blog/index")
		}
	} else {
		htmlHandler(c, "注册失败!账号已被注册!", "/blog/admin/login")
	}
}

func htmlHandler(c *gin.Context, content, redirectURL string) {
	// redirectURL使用相对路径而不是绝对URL
	c.HTML(http.StatusOK, "endLogin.tmpl", gin.H{
		"redirectURL": redirectURL,
		"content":     content,
	})
}
func SignInHandler(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	result := repository.CheckUser(id)
	if result {
		user, err := repository.GetUser(id)
		if err != nil {
			htmlHandler(c, "数据库初始化失败!", "/blog/admin/login")
		} else {
			result, err = utils.VerifyPassword(password, user.Password)
			if err != nil {
				htmlHandler(c, "密码错误!", "/blog/admin/login")
			} else {
				htmlHandler(c, "登录成功!", "/blog/index")
			}
		}
	} else {
		htmlHandler(c, "账号不存在!", "/blog/admin/login")
	}
}
