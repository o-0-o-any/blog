package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func SessionConfig() cookie.Store {
	// 创建Cookie存储引擎
	store := cookie.NewStore([]byte("secret-key-aaa2021s"))

	// Session存储规则的设置
	store.Options(sessions.Options{
		MaxAge:   3600, // 1h Cookie过期时间
		Path:     "/",  // Cookie的生效地址 /表示全部网站
		HttpOnly: true, // 防止XXS攻击
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // 防止CSRF攻击
	})

	return store
}
