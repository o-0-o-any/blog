package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionData(c *gin.Context) string {
	session := sessions.Default(c)
	// 获取当前登录的账号信息
	id := session.Get("id")

	// 直接将string类型的id返回 后续使用id查询数据库时有类型转换功能
	// id是interface{}类型 进行类型断言
	if str, ok := id.(string); ok {
		return str
	}

	return ""
}
