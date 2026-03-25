package controller

import (
	"blog/models"
	"blog/repository"
	"blog/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
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
				// 登录成功 将数据以键值对形式写入Session中
				// 得到当前请求的Session对象
				session := sessions.Default(c)
				// 数据存储
				session.Set("id", id)
				if err := session.Save(); err != nil {
					htmlHandler(c, "登录失败!", "/blog/index")
				} else {
					htmlHandler(c, "登录成功!", "/blog/index")
				}
			}
		}
	} else {
		htmlHandler(c, "账号不存在!", "/blog/admin/login")
	}
}

func SignOutHandler(c *gin.Context) {
	// 退出登录
	// 可以执行这个函数的 都是已经登录过的 所以Cookie都不是空的
	// 需要退出登录的话 就需要先清空Cookie
	c.SetCookie(
		"blog-session", // Cookie全称
		"",             // 值清空
		-1,             // 过期时间设置为-1 即立即删除
		"/",            // 与前面保持一致
		"",             // 域名
		false,          // secure 是否仅HTTPS
		true,           // 设置为true 防止js读取
	)

	// 然后重定向到当前界面的无登录的layout模板界面了  渲染延时重定向界面  防止卡住
	htmlHandler(c, "正在退出中...", "/blog/index")
}
