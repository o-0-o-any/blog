package controller

import (
	"blog/models"
	"blog/repository"
	"blog/utils"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func dashBoardHandler(templateFile string) func(c *gin.Context, data any) {
	// template解析渲染组合模板
	return func(c *gin.Context, data any) {
		// 获取Session中保存的id
		session := sessions.Default(c)
		id := session.Get("id")

		// 将id与data融合 一起传入渲染的数据中 "id": id 用来作为url中id参数
		allData := gin.H{
			"id": id,
		}
		if data != nil {
			if d, ok := data.(gin.H); ok {
				for key, value := range d {
					allData[key] = value
				}
			}
		}

		// 解析模板文件
		t, err := template.ParseFiles("templates/dashboard/layout.tmpl", templateFile)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}

		// 渲染模板文件  传入加入了id键值对的数据
		if err := t.ExecuteTemplate(c.Writer, "layout", allData); err != nil {
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

	// 根据编辑URL中的id参数 先查找用户信息 检查用户有没有完善必要信息!
	id := c.Param("id")
	user, err := repository.GetUser(id)
	if err != nil {
		redirectHandler(c, "添加博客失败!账号错误!请联系管理员!", "/blog/admin/"+id+"/dashboard")
	} else {
		if user.Name == "暂未设置" { // 如果姓名没有完善 则需要先完善个人信息
			redirectHandler(c, "添加博客失败!请先完善个人信息!", "/blog/admin/"+id+"/dashboard/info")
		}
	}

	// 获取前段数据
	article := models.ArticlesModel{}

	article.Title = c.PostForm("title")
	article.Type = c.PostForm("author")
	article.Text = c.PostForm("text")
	// 生成time数据
	article.Date = time.Now()

	// 插入数据库中
	if err := repository.AddArticles(article); err != nil {
		redirectHandler(c, "添加博客失败!请联系管理员!", "/blog/admin/"+id+"/dashboard")
	} else {
		redirectHandler(c, "添加成功!", "/blog/admin/"+id+"/dashboard")
	}
}

// ===== 个人信息界面 =====
func InfoHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/info.tmpl")(c, nil)
}

// ===== 修改个人信息界面
func EditInfoHTMLHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/infoEdit.tmpl")(c, nil)
}
func EditInfoHandler(c *gin.Context) {
	user := models.UserModel{}
	user.Name = c.PostForm("name")
	user.Age, _ = utils.StringTOInt(c.PostForm("Age")) // 前端只能输入数字 直接把error省略 也可能是空字符串
	user.Gender = c.PostForm("Gender")
	user.Email = c.PostForm("Email")

	// 这里不知道用啥方法了 反正就是判断至少有一个表单不是空的

	// 延迟跳转到指定URL
	redirectHandler(c, "修改成功!", "/blog/admin/"+c.Param("id")+"/dashboard/info")
}
