package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	// 加载模板文件
	r.LoadHTMLFiles(
		"templates/login/login.tmpl",
		"templates/login/endLogin.tmpl",

		"templates/index/layout.tmpl",
		"templates/index/index.tmpl",
		"templates/index/message.tmpl",
		"templates/index/driver.tmpl",

		"templates/message/edit.tmpl",
	)

	// 路由注册 路由组
	allGroup := r.Group("/blog")
	{
		// ===博客主页===
		indexGroup := allGroup.Group("/index")
		{
			// 博客主页
			indexGroup.GET("/", controller.IndexHandler)
			// 文章列表
			indexGroup.GET("/message", controller.MessageHandler)
			// 导航页
			indexGroup.GET("/driver", controller.DriverHandler)
		}

		// ===用户登录===
		adminGroup := allGroup.Group("/admin")
		{
			// 用户主页
			adminGroup.GET("/index", controller.AdminIndexHandler)
			// 登录注册功能
			loginGroup := adminGroup.Group("/login")
			{
				// 登录界面
				loginGroup.GET("/", controller.LoginHandler)
				// 注册
				loginGroup.POST("/signUp", controller.SignUpHandler)
				// 登录
				loginGroup.POST("/signIn", controller.SignInHandler)
			}
		}

		// ===博客内容===
		blogGroup := allGroup.Group("/message")
		{
			// 博客编辑
			blogGroup.GET("/edit", controller.EditHandler)

			// 博客网页
			blogGroup.GET("/:id", controller.PerMessageHandler)
		}

	}

	return r
}
