package controller

import (
	"blog/models"
	"blog/repository"
	"blog/utils"
	"fmt"
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
	data := gin.H{}
	// 用来渲染当前用户发布的全部博客内容
	userID := utils.GetSessionData(c)
	user, _ := repository.GetUser(userID) // 登录之后即存在这个账号 省略掉error
	// 根据userName去寻找articles
	articles, err := repository.FindUserAllArticle(user.Name)
	if err != nil {
		data["content"] = "系统异常!请联系管理员!"
	} else {
		data["content"] = "你发布的博客"
		// 设置博客跳转的RedirectURL
		for i := 0; i < len(articles); i++ {
			articles[i].RedirectURL = fmt.Sprintf("/blog/admin/%s/dashboard/article/%d", user.ID, articles[i].Id)
		}
		// 将articles放入gin.H{}中
		data["articles"] = articles
	}
	dashBoardHandler("templates/dashboard/index.tmpl")(c, data)
}

// ===== 关于博客编辑 =====
func EditHandler(c *gin.Context) {
	id := utils.GetSessionData(c)
	dashBoardHandler("templates/dashboard/edit.tmpl")(c, gin.H{
		"id": id,
	})
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
	article.Type = c.PostForm("type")
	article.Text = c.PostForm("text")
	article.Author = user.Name
	// 生成time数据
	article.Date = time.Now()

	// 插入数据库中
	if err := repository.AddArticles(article); err != nil {
		redirectHandler(c, "添加博客失败!请联系管理员!", "/blog/admin/"+id+"/dashboard")
	} else {
		redirectHandler(c, "添加成功!", "/blog/admin/"+id+"/dashboard")
	}
}

// ====== 获取当前用户信息 =====
func getCurrentUserInfo(c *gin.Context) gin.H {
	// 得到当前登录的用户信息
	id := utils.GetSessionData(c)
	user, err := repository.GetUser(id)
	data := gin.H{}
	if err != nil {
		data = nil
	} else {
		data = gin.H{
			"id":         user.ID,
			"age":        user.Age,
			"name":       user.Name,
			"gender":     user.Gender,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
		}
	}
	return data
}

// ===== 个人信息界面 =====
func InfoHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/info.tmpl")(c, getCurrentUserInfo(c))
}

// ===== 修改个人信息界面
func EditInfoHTMLHandler(c *gin.Context) {
	dashBoardHandler("templates/dashboard/infoEdit.tmpl")(c, getCurrentUserInfo(c))
}
func EditInfoHandler(c *gin.Context) {
	user := models.UserModel{}
	// 获取用户id
	user.ID = utils.GetSessionData(c) // user.ID是string类型

	// 得到用户修改前的信息
	beforeUser, _ := repository.GetUser(user.ID) // 对登录的用户暂时忽略掉error
	fmt.Println("beforeUser:", beforeUser)
	// 更新数据
	user.Name = c.PostForm("name")
	// 对姓名进行查重
	if user.Name != beforeUser.Name && repository.JudgeNameisExit(user.Name) {
		redirectHandler(c, "姓名已经存在!", "/blog/admin/"+c.Param("id")+"/dashboard/info")
		return
	}
	user.Age, _ = utils.StringTOInt(c.PostForm("age")) // 前端只能输入数字 直接把error省略 也可能是空字符串
	user.Gender = c.PostForm("gender")
	user.Email = c.PostForm("email")

	// 这里不知道用啥方法了 反正就是判断至少有一个表单不是空的 暂时空住

	// 更新数据库数据 这有一点要注意的是 个人信息变化后 这个作者写过的博客中的作者信息也要变化
	// 先更新文章的作者信息 然后再更新作者信息 不然会截断操作
	arrticles, _ := repository.FindUserAllArticle(beforeUser.Name) // 使用修改前的名字查询 默认进行 暂时想不出来怎么对错误提示
	// 对前后名字判断 如果变化则修改
	if beforeUser.Name != user.Name {
		for i := 0; i < len(arrticles); i++ {
			// 修改作者信息
			arrticles[i].Author = user.Name
			// 执行更新操作
			_ = repository.UpdateArticleInfo(arrticles[i]) // 默认执行修改博客信息 暂时不知道怎么去处理异常错误
		}
	}

	err1 := repository.UpdateUser(user)
	// 延迟跳转到指定URL
	if err1 != nil {
		redirectHandler(c, "修改出现问题!请联系管理员!", "/blog/admin/"+c.Param("id")+"/dashboard/info")
	} else {
		redirectHandler(c, "修改成功!", "/blog/admin/"+c.Param("id")+"/dashboard/info")
	}
}
