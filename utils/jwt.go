package utils

import (
	"blog/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token密钥
var TokenSecret = []byte("AnY-blog-secret")

// 单Token认证
// 生成Token
func GenerateJWT(id string, name string) (string, error) {
	claim := models.Claim{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 设置Token过期时间为72小时
			IssuedAt:  time.Now().Unix(),                     // 签发时间
			Issuer:    "史耀宇",                                 // 签发者
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(TokenSecret)
}

// 解析Token
func ParseToken(tokenStr string) (*models.Claim, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&models.Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("blog"), nil
		})

	// 判断Token是否有效
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token is invalid %v", err)
}
