package controller

import (
	"blog/repository"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	LayoutSignOut = "templates/index/layoutSignOut.tmpl" // 未登录的layout文件路径
	LayoutSignIn  = "templates/index/layoutSignIn.tmpl"  // 登录的layout文件路径
)

func handler(templateFile string) func(c *gin.Context, data any) {
	return func(c *gin.Context, data any) {
		var layoutFile string
		// ===根据Session判断是否登录 然后决定使用哪个layout模板===
		session := sessions.Default(c)
		id := session.Get("id")
		if id != nil {
			// 登录状态
			layoutFile = LayoutSignIn
		} else {
			layoutFile = LayoutSignOut
		}

		// ===解析组合模板===
		t, err := template.ParseFiles(layoutFile, templateFile)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}

		// ===对数据处理 在data基础上新增JSON数据===
		allData := gin.H{
			"name": id,
		}
		if data != nil {
			if d, ok := data.(gin.H); ok {
				for key, value := range d {
					allData[key] = value
				}
			}
		}

		// ===渲染组合模板===
		err = t.ExecuteTemplate(c.Writer, "layout", allData)
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

// uint类型转换成string类型
func intTURNstring(u int) string {
	return strconv.Itoa(int(u))
}

func MessageHandler(c *gin.Context) {
	// 从数据库中得到全部博客信息
	articles, err := repository.FindAllArticles()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"content": "数据初始化失败!请联系管理员!",
		})
	}
	// 由于刚提交博客时没有给RedirectURL赋值 因此在这里指定RedirectURL与string(id)保持一致
	for i := 0; i < len(articles); i++ {
		articles[i].RedirectURL = "/blog/message/article/" + intTURNstring(articles[i].Id)
	}
	handler("templates/index/message.tmpl")(c, gin.H{
		"articles": articles,
	})
}

func DriverHandler(c *gin.Context) {
	handler("templates/index/driver.tmpl")(c, nil)
}
