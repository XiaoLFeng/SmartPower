package auth

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/dto/duser"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/util"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// AuthLogin
//
// # 用户登录逻辑
//
// 用户登录逻辑, 登录成功后返回用户信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - user: string, 用户名/邮箱/手机号
//   - password: string, 密码
//
// # 返回:
//   - dUser: *duser.UserCurrent, 用户信息
//   - err: error, 错误
func (s *sAuth) AuthLogin(ctx context.Context, user string, password string) (dUser *duser.UserCurrent, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 AuthLogin | 用户登录")
	// 检查用户是否存在
	var getUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Username: user}).
		WhereOr(do.XfUser{Phone: user}).
		WhereOr(do.XfUser{Email: user}).
		Scan(&getUser)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getUser == nil {
		return nil, berror.NewError(bcode.NotExist, "用户 "+user+" 不存在")
	}
	// 对用户进行密码验证
	if util.VerifyPassword(getUser.Password, password) {
		// 生成令牌
		tUUID, _ := uuid.NewV7()
		tokenUUID, _ := uuid.NewV7()
		_, err = dao.XfToken.Ctx(ctx).Data(do.XfToken{
			Tuuid:     tUUID.String(),
			Uuid:      getUser.Uuid,
			Token:     tokenUUID.String(),
			ExpiredAt: gtime.NewFromTimeStamp(gtime.Timestamp() + 43200),
		}).Insert()
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		} else {
			returnUser := duser.UserCurrent{
				User: duser.DUser{
					Uuid:     getUser.Uuid,
					Username: getUser.Username,
					Email:    getUser.Email,
					Phone:    getUser.Phone,
					Role:     getUser.Role,
				},
				Token: tokenUUID.String(),
			}

			// 对数据进行整合
			return &returnUser, nil
		}
	} else {
		return nil, berror.NewError(bcode.OperationFailed, "密码错误")
	}
}
