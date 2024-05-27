package user

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/dto/dcompany"
	"SmartPower/internal/model/dto/duser"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// GetUserCurrent
//
// # 获取当前用户信息
//
// 获取当前用户信息, 用于获取当前登录用户的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - getUser: *duser.UserDetailed, 获取当前用户信息返回
//   - err: error, 错误信息
func (s *sUser) GetUserCurrent(ctx context.Context) (getUser *duser.UserDetailed, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 GetUserCurrent | 获取当前用户信息")
	// 获取请求头
	getRequest := g.RequestFromCtx(ctx)
	// 获取用户 UUID
	getToken := getRequest.Header.Get("Authorization")
	// 获取用户信息
	userInfo, err := service.Token().GetUserByToken(ctx, getToken)
	if err != nil {
		return nil, err
	} else {
		// 根据用户信息获取企业
		var getCompany *entity.XfCompanies
		err := dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: userInfo.Uuid}).Scan(&getCompany)
		if err != nil {
			return nil, err
		} else {
			getUser := &duser.UserDetailed{
				User: duser.DUser{
					Uuid:     userInfo.Uuid,
					Username: userInfo.Username,
					Email:    userInfo.Email,
					Phone:    userInfo.Phone,
					Role:     userInfo.Role,
				},
			}
			if getCompany != nil {
				getUser.Company = dcompany.DCompany{
					Cods:           getCompany.Cods,
					Name:           getCompany.Name,
					Address:        getCompany.Address,
					Representative: getCompany.Representative,
				}
			}
			return getUser, nil
		}
	}
}
