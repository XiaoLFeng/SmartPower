package auth

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/util"
	"context"
	"github.com/gogf/gf/v2/os/glog"
)

// ChangePassword
//
// # 修改密码
//
// 修改密码接口, 用于用户修改密码. 用户在这里可以进行修改密码，修改密码的时候需要用户提供当前密码以及新密码，修改成功后会返回修改成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - getUser: entity.XfUser, 用户信息
//   - nowPassword: string, 当前密码
//   - newPassword: string, 新密码
//
// # 返回:
//   - err: error, 错误
func (s *sAuth) ChangePassword(ctx context.Context, getUser entity.XfUser, nowPassword string, newPassword string) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 ChangePassword | 修改密码")
	// 检查密码是否正确
	if util.VerifyPassword(getUser.Password, nowPassword) {
		// 修改密码
		getUser.Password = util.EncryptPassword(newPassword)
		// 更新密码
		_, err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: getUser.Uuid}).Update(getUser)
		if err != nil {
			return xerror.NewErrorHasError(xerror.ServerInternalError, err)
		} else {
			return nil
		}
	} else {
		return xerror.NewError(xerror.OperationFailed, "密码错误")
	}
}
