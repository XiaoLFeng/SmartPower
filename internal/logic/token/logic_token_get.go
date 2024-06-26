package token

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetUserByToken
//
// # 根据令牌获取用户信息
//
// 根据令牌获取用户信息, 传入令牌进行检查, 若令牌存在则返回用户信息. 对输入的令牌进行检查，只允许搜索挡枪用户存在的令牌. 不允许跨用户查询其他用户
// 的令牌相关的信息. 若令牌不存在则返回对应的报错信息，若令牌存在且为正确的内容则返回用户信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - token: string, 令牌
//
// # 返回:
//   - user: *entity.XfUser, 用户信息
//   - err: error, 错误
func (s *sToken) GetUserByToken(ctx context.Context, token string) (user *entity.XfUser, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 GetUserByToken | 根据令牌获取用户信息")
	var getToken *entity.XfToken
	err = dao.XfToken.Ctx(ctx).Where(do.XfToken{Token: token}).Scan(&getToken)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getToken == nil {
		return nil, berror.NewError(bcode.NotExist, "令牌不存在")
	}
	// 令牌存在，检查是否过期
	if getToken.ExpiredAt.After(gtime.Now()) {
		// 获取用户信息
		var getUser *entity.XfUser
		err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: getToken.Uuid}).Scan(&getUser)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		if getUser == nil {
			return nil, berror.NewError(bcode.NotExist, "用户不存在")
		} else {
			return getUser, nil
		}
	} else {
		// 令牌过期并且删除内容
		_, _ = dao.XfToken.Ctx(ctx).Where(do.XfToken{Token: token}).Delete()
		return nil, berror.NewError(bcode.Expired, "令牌已过期")
	}
}
