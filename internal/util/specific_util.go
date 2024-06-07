package util

import (
	"crypto/md5"
	"github.com/google/uuid"
)

// GetAdminRoleUUID
//
// # 获取管理员角色UUID
//
// 用于获取管理员角色UUID, 用于创建超级管理员时使用. 以及其他需要使用管理员角色UUID的地方.
//
// # 返回:
//   - string: 管理员角色UUID
func GetAdminRoleUUID() string {
	return uuid.UUID(md5.Sum([]byte("admin"))).String()
}

// GetUserRoleUUID
//
// # 获取用户角色UUID
//
// 用于获取用户角色UUID, 用于创建用户时使用. 以及其他需要使用用户角色UUID的地方.
//
// # 返回:
//   - string: 用户角色UUID
func GetUserRoleUUID() string {
	return uuid.UUID(md5.Sum([]byte("user"))).String()
}
