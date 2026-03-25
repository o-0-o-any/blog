package controller

import (
	"blog/models"
	"blog/repository"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func dashBoardHandler(templateFile string) func(c *gin.Context, data any) {
	// template解析渲染组合模板
	return func(c *gin.Context, data any) {
		t, err := template.ParseFiles("templates/dashboard/layout.tmpl", templateFile)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}

		// 渲染模板文件
		if err := t.ExecuteTemplate(c.Writer, "layout", data); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	}
}

// ===== 关于仪表盘主页 =====
func DashboardIndexHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/index.tmpl")(c, nil)
}

// ===== 关于博客编辑 =====
func EditHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/edit.tmpl")(c, nil)
}
func PostEditHandler(c *gin.Context) {
	// 推送博客的post操作 主要进行获取请求数据与数据库存储
	// 获取前段数据
	article := models.ArticlesModel{}
	article.Title = c.PostForm("title")
	article.Author = c.PostForm("author")
	article.Text = c.PostForm("text")
	// 生成time数据
	article.Date = time.Now()

	// 插入数据库中
	if err := repository.AddArticles(article); err != nil {
		redirectHandler(c, "添加博客失败!请联系管理员!", "/blog/admin/dashboard")
	} else {
		redirectHandler(c, "添加成功!", "/blog/admin/dashboard")
	}
}

// ===== 个人信息界面 =====
func InfoHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/info.tmpl")(c, nil)
}
