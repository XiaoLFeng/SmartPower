package util

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

// EncodePassword
//
// # 加密密码
//
// 加密密码, 用于对密码进行加密处理. 加密密码的时候会对密码进行base64编码处理, 然后再对密码进行bcrypt加密处理.
//
// # 参数:
//   - password: string, 密码
func EncodePassword(password string) string {
	getBase6Password := base64.StdEncoding.EncodeToString([]byte(password))
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(getBase6Password), bcrypt.DefaultCost)
	return string(fromPassword)
}
