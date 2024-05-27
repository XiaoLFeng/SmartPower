package auth

import (
	v1 "SmartPower/api/auth/v1"
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/util"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// UserRegister
//
// # 用户注册
//
// 用户注册接口, 用于用户注册. 用户在这里可以进行注册，注册的时候需要用户提供基本注册信息以及企业的认证信息，注册成功后会返回注册成功的信息.
// 用户的注册需要用户提供必要注册相关的信息，注册成功后会返回注册成功的信息. 若注册失败将会抛出指定的错误信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.AuthRegisterReq, 请求
//
// # 返回:
//   - err: error, 错误
func (s *sAuth) UserRegister(ctx context.Context, req *v1.AuthRegisterReq) (err error) {
	// 对输入信息进行检查
	var getUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).
		Where(do.XfUser{Username: req.Username}).
		WhereOr(do.XfUser{Phone: req.Phone}).
		WhereOr(do.XfUser{Email: req.Email}).
		Limit(1).
		Scan(&getUser)
	if getUser != nil {
		return xerror.NewErrorHasError(xerror.UserAlreadyExists, err)
	}
	// 检查企业是否存在
	var getEnterprise *entity.XfCompanies
	err = dao.XfCompanies.Ctx(ctx).
		Where(do.XfCompanies{Name: req.Company}).
		WhereOr(do.XfCompanies{Cods: req.CompanyCods}).
		Limit(1).
		Scan(&getEnterprise)
	if getEnterprise != nil {
		return xerror.NewErrorHasError(
			xerror.BaseLocalCode("CompanyAlreadyExists", 201, "企业已存在"),
			err,
		)
	}
	// 创建用户
	err = dao.XfUser.Transaction(ctx, func(_ context.Context, tx gdb.TX) error {
		var hasError error
		getUuid, _ := uuid.NewV7()
		// 创建用户
		_, hasError = tx.Ctx(ctx).Insert(dao.XfUser.Table(), do.XfUser{
			Uuid:     getUuid.String(),
			Username: req.Username,
			Email:    req.Email,
			Phone:    req.Phone,
			Password: util.EncodePassword(req.Password),
		})
		// 企业注册
		_, hasError = tx.Ctx(ctx).Insert(dao.XfCompanies.Table(), do.XfCompanies{
			Meta:           g.Meta{},
			Cods:           req.CompanyCods,
			Name:           req.Company,
			Address:        req.CompanyAddress,
			Representative: req.Representative,
			Uuid:           getUuid.String(),
		})
		return hasError
	})
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	return nil
}
