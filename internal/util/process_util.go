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
//
// # 返回:
//   - string, 加密后的密码
func EncodePassword(password string) string {
	getBase64Password := base64.StdEncoding.EncodeToString([]byte(password))
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(getBase64Password), bcrypt.DefaultCost)
	return string(fromPassword)
}

// VerifyPassword
//
// # 验证密码
//
// 验证密码, 用于对密码进行验证处理. 验证密码的时候会对密码进行base64编码处理, 然后再对密码进行bcrypt验证处理.
//
// # 参数:
//   - hashPassword: string, 加密后的密码
//   - password: string, 密码
//
// # 返回:
//   - bool, 是否验证成功
func VerifyPassword(hashPassword string, password string) bool {
	getBase64Password := base64.StdEncoding.EncodeToString([]byte(password))
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(getBase64Password))
	return err == nil
}
