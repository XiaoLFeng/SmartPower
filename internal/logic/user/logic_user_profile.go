package user

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// UserEditProfile
//
// # 编辑用户信息
//
// 编辑用户信息. 用户在这里可以进行编辑用户信息，编辑用户信息的时候需要用户提供邮箱, 手机号，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - phone: string, 手机号
//   - email: string, 邮箱
//
// # 返回:
//   - err: error, 错误信息
func (s *sUser) UserEditProfile(ctx context.Context, phone string, email string) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 UserEditProfile | 编辑用户信息")
	// 检查用户是否存在
	getRequest := g.RequestFromCtx(ctx)
	getToken := getRequest.Header.Get("Authorization")
	userInfo, err := service.Token().GetUserByToken(ctx, getToken)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if userInfo == nil {
		return xerror.NewError(xerror.OperationFailed, "用户不存在")
	}
	// 更新用户信息
	_, err = dao.XfUser.Ctx(ctx).
		Data(do.XfUser{Email: email, Phone: phone}).
		Where(do.XfUser{Uuid: userInfo.Uuid}).
		Update()
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	return nil
}
