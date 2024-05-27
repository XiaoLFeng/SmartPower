package duser

// UserCurrent
//
// # 用户信息
//
// 用于获取用户的一些基本信息，这些信息用于数据之间的传输工作使用
//
// # 参数
//   - User: DUser, 用户信息
//   - Token: string 返回授权的 token
type UserCurrent struct {
	User  DUser  `json:"user"`
	Token string `json:"token"`
}

// DUser
//
// # 用户基本信息
//
// 用于返回用户从数据库获取的信息后进行数据脱敏的操作
//
// # 参数
//   - Uuid: string, 用户的 UUID
//   - Username: string, 用户的用户名
//   - Email: string, 用户的邮箱
//   - Phone: string, 用户的手机号
//   - Role: string, 角色，不为角色的 UUID，是角色的 Name
type DUser struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}
