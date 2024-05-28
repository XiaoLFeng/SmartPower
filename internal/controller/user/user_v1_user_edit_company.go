package user

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/user/v1"
)

// UserEditCompany
//
// # 编辑企业信息
//
// 编辑企业信息. 用户在这里可以进行编辑企业信息，编辑企业信息的时候需要用户提供企业编码, 企业名称, 企业地址, 企业法人，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.UserEditCompanyReq, 编辑企业信息请求
//
// # 返回:
//   - res: *v1.UserEditCompanyRes, 编辑企业信息返回
//   - err: error, 错误信息
func (c *ControllerV1) UserEditCompany(ctx context.Context, req *v1.UserEditCompanyReq) (res *v1.UserEditCompanyRes, err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 UserEditCompany | 编辑企业信息")
	err = service.User().UserEditCompany(ctx, req.Cods, req.Name, req.Address, req.Representative)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
