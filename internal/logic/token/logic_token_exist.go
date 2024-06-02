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

// CheckTokenExist
//
// # 检查令牌是否存在
//
// 检查令牌是否存在, 传入令牌进行检查, 若令牌存在则返回错误信息. 对输入的令牌进行检查，只允许搜索挡枪用户存在的令牌. 不允许跨用户查询其他用户
// 的令牌相关的信息. 若令牌不存在则返回对应的报错信息，若令牌存在且为正确的内容则不进行任何报错的抛出.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - token: string, 令牌
//
// # 返回:
//   - err: error, 错误
func (s *sToken) CheckTokenExist(ctx context.Context, token string) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 CheckTokenExist | 检查令牌是否存在")
	// 对 token 进行检查
	var getToken *entity.XfToken
	err = dao.XfToken.Ctx(ctx).Where(do.XfToken{Token: token}).Scan(&getToken)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getToken == nil {
		return berror.NewError(bcode.NotExist, "令牌不存在")
	}
	// 令牌存在，检查是否过期
	if getToken.ExpiredAt.Before(gtime.Now()) {
		// 令牌过期并且删除内容
		_, _ = dao.XfToken.Ctx(ctx).Where(do.XfToken{Token: token}).Delete()
		return berror.NewError(bcode.Expired, "令牌已过期")
	} else {
		return nil
	}
}
