package routers

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	// 加载模板文件
	r.LoadHTMLFiles(
		"templates/login/login.tmpl",
		"templates/login/endLogin.tmpl",

		"templates/index/layoutSignOut.tmpl",
		"templates/index/layoutSignIn.tmpl",
		"templates/index/index.tmpl",
		"templates/index/message.tmpl",
		"templates/index/driver.tmpl",

		"templates/message/endEdit.tmpl",
		"templates/message/article.tmpl",

		"templates/dashboard/layout.tmpl",
		"templates/dashboard/edit.tmpl",
		"templates/dashboard/info.tmpl",
		"templates/dashboard/infoEdit.tmpl",
	)

	// 注册Session中间件到全部路由中
	r.Use(sessions.Sessions("blog-session", middleware.SessionConfig()))

	// 路由注册 路由组
	allGroup := r.Group("/blog")
	{
		// ===博客主页===
		// 首页
		allGroup.GET("/index", controller.IndexHandler)
		// 文章列表
		allGroup.GET("/message", controller.MessageHandler)
		// 导航页
		allGroup.GET("/driver", controller.DriverHandler)

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
				// 退出登录
				loginGroup.POST("/signOut", controller.SignOutHandler)
			}
			// 每个用户的仪表盘
			// 仪表盘
			dashboardGroup := adminGroup.Group("/:id/dashboard")
			{
				// 仪表盘主页
				dashboardGroup.GET("/", controller.DashboardIndexHandler)
				dashboardGroup.GET("/index", controller.DashboardIndexHandler)
				// 博客编辑
				dashboardGroup.GET("/edit", controller.EditHandler)
				dashboardGroup.POST("/edit", controller.PostEditHandler)
				// 个人信息
				dashboardGroup.GET("/info", controller.InfoHandler)
				dashboardGroup.GET("/info/edit", controller.EditInfoHTMLHandler)
				dashboardGroup.POST("/info/edit", controller.EditInfoHandler)
			}
		}

		// ===博客内容===
		blogGroup := allGroup.Group("/message")
		{
			// 博客网页
			blogGroup.GET("/article/:id", controller.PerMessageHandler)
		}

	}

	return r
}
