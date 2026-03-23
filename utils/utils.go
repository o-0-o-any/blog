package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword对密码进行哈希加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func VerifyPassword(password, hashPwd string) (bool, error) {
	// bcrypt.CompareHashAndPassword判断密码的正确 分别传入哈希加密的密码与需要对比的密码
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))

	if err != nil {
		// 密码错误的情况
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		// 其他情况 密码验证失败
		return false, fmt.Errorf("error comparing passwords: %v", err)
	}
	// 密码正确
	return true, nil

}
