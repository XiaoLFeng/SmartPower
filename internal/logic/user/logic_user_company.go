package user

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
)

// UserEditCompany
//
// # 编辑企业信息
//
// 编辑企业信息. 用户在这里可以进行编辑企业信息，编辑企业信息的时候需要用户提供企业编码, 企业名称, 企业地址, 企业法人，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - cods: string, 企业编码
//   - name: string, 企业名称
//   - address: string, 企业地址
//   - principal: string, 企业法人
//
// # 返回:
//   - err: error, 错误信息
func (s *sUser) UserEditCompany(
	ctx context.Context,
	cods string,
	name string,
	address string,
	principal string,
) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 UserEditCompany | 编辑企业信息")
	// 获取用户登陆所在的企业
	company, err := service.Electric().GetCompanyByHeader(ctx)
	if err != nil {
		return err
	}
	// 更新企业信息
	_, err = dao.XfCompanies.Ctx(ctx).
		Data(do.XfCompanies{Cods: cods, Name: name, Address: address, Representative: principal}).
		Where(do.XfCompanies{Uuid: company.Uuid, Cods: company.Cods}).
		Update()
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	return nil
}
