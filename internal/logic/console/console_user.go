package console

import (
	v1 "SmartPower/api/console/v1"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/dto/dcompany"
	"SmartPower/internal/model/dto/duser"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/util"
	"context"
	"crypto/md5"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// UserAdd
//
// # 添加控制台用户
//
// 添加控制台用户, 用于添加控制台用户. 传入用户名, 密码, 角色.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - getData: *v1.ConsoleUserAddReq, 添加控制台用户请求
//
// # 返回:
//   - err: error, 错误
func (s *sConsole) UserAdd(ctx context.Context, getData *v1.ConsoleUserAddReq) (user *duser.DUserForConsoleCreate, err error) {
	g.Log().Noticef(ctx, "[LOGIC] 执行 UserAdd | 添加控制台用户")
	// 检查信息是否完备化
	if !getData.IsAdmin {
		if getData.Company == "" || getData.CompanyCods == "" || getData.CompanyAddress == "" || getData.Representative == "" {
			return nil, berror.NewError(bcode.OperationFailed, "公司信息不完整")
		}
	}
	// 检查用户是否已存在
	var checkUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).
		Where(do.XfUser{Email: getData.Email}).
		WhereOr(do.XfUser{Username: getData.Username}).
		WhereOr(do.XfUser{Phone: getData.Phone}).
		Scan(&checkUser)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if checkUser != nil {
		return nil, berror.NewError(bcode.OperationFailed, "用户已存在")
	}
	// 创建用户
	generatePassword := util.GenerateRandomString(10)
	getUUID, err := uuid.NewV7()
	var getRole string
	var roleName string
	if getData.IsAdmin {
		getRole = util.GetAdminRoleUUID()
		roleName = "admin"
	} else {
		getRole = util.GetUserRoleUUID()
		roleName = "user"
	}
	_, err = dao.XfUser.Ctx(ctx).Data(do.XfUser{
		Uuid:     getUUID.String(),
		Username: getData.Username,
		Email:    getData.Email,
		Phone:    getData.Phone,
		Password: util.EncryptPassword(generatePassword),
		Role:     getRole,
	}).Insert()
	// 如果不是管理员需要创建公司
	if !getData.IsAdmin {
		_, err = dao.XfCompanies.Ctx(ctx).Data(do.XfCompanies{
			Cods:           getData.CompanyCods,
			Name:           getData.Company,
			Address:        getData.CompanyAddress,
			Representative: getData.Representative,
			Uuid:           getUUID.String(),
		}).Insert()
	}
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return &duser.DUserForConsoleCreate{
		User: duser.DUser{
			Uuid:     getUUID.String(),
			Username: getData.Username,
			Email:    getData.Email,
			Phone:    getData.Phone,
			Role:     roleName,
		},
		Password: generatePassword,
	}, nil
}

// UserResetPassword
//
// # 重置控制台用户密码
//
// 重置控制台用户密码, 用于重置控制台用户密码. 传入用户ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: v1.ConsoleUserResetPasswordReq, 重置控制台用户密码请求
//
// # 返回:
//   - err: error, 错误
func (s *sConsole) UserResetPassword(ctx context.Context, req *v1.ConsoleUserResetPasswordReq) (User *duser.DUserForConsoleCreate, err error) {
	g.Log().Noticef(ctx, "[LOGIC] 执行 UserResetPassword | 重置控制台用户密码")
	// 获取用户
	var getUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: req.UUID}).Scan(&getUser)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getUser == nil {
		return nil, berror.NewError(bcode.OperationFailed, "用户不存在")
	}
	// 重置密码
	generatePassword := util.GenerateRandomString(10)
	_, err = dao.XfUser.Ctx(ctx).Data(do.XfUser{Password: util.EncryptPassword(generatePassword)}).Where(do.XfUser{Uuid: req.UUID}).Update()
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 根据角色获取角色名
	var getRole *entity.XfRole
	err = dao.XfRole.Ctx(ctx).Where(do.XfRole{Ruuid: getUser.Role}).Scan(&getRole)
	return &duser.DUserForConsoleCreate{
		User: duser.DUser{
			Uuid:     getUser.Uuid,
			Username: getUser.Username,
			Email:    getUser.Email,
			Phone:    getUser.Phone,
			Role:     getRole.Name,
		},
		Password: generatePassword,
	}, nil
}

// UserDelete
//
// # 删除控制台用户
//
// 删除控制台用户, 用于删除控制台用户. 传入用户ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - getUUID: string, 用户ID
//
// # 返回:
//   - err: error, 错误
func (s *sConsole) UserDelete(ctx context.Context, getUUID string) (err error) {
	g.Log().Noticef(ctx, "[LOGIC] 执行 UserDelete | 删除控制台用户")
	// 检查是否为超级管理员
	if getUUID == uuid.UUID(md5.Sum([]byte("SmartPowerSuperConsoleUser"))).String() {
		return berror.NewError(bcode.OperationFailed, "不能删除超级管理员")
	}
	// 检查用户是否存在
	var getUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: getUUID}).Scan(&getUser)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getUser == nil {
		return berror.NewError(bcode.OperationFailed, "用户不存在")
	}
	// 删除用户
	_, err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: getUUID}).Delete()
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}

// UserEdit
//
// # 编辑控制台用户
//
// 编辑控制台用户, 用于编辑控制台用户. 传入用户ID, 用户名, 邮箱, 手机号, 公司名称, 公司编码, 公司地址, 法人代表, 是否为管理员.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - getData: *v1.ConsoleUserEditReq, 编辑控制台用户请求
//
// # 返回:
//   - err: error, 错误
func (s *sConsole) UserEdit(ctx context.Context, getData *v1.ConsoleUserEditReq) (err error) {
	g.Log().Noticef(ctx, "[LOGIC] 执行 UserEdit | 编辑控制台用户")
	// 检查信息是否完备化
	if !getData.IsAdmin {
		if getData.Company == "" || getData.CompanyCods == "" || getData.CompanyAddress == "" || getData.Representative == "" {
			return berror.NewError(bcode.OperationFailed, "公司信息不完整")
		}
	}
	// 检查用户是否存在
	var getUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: getData.UUID}).Scan(&getUser)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getUser == nil {
		return berror.NewError(bcode.OperationFailed, "用户不存在")
	}
	// 检查主要编辑后的信息是否已被注册
	var checkUser *entity.XfUser
	err = dao.XfUser.Ctx(ctx).
		Where(do.XfUser{Email: getData.Email}).
		WhereOr(do.XfUser{Username: getData.Username}).
		WhereOr(do.XfUser{Phone: getData.Phone}).
		Scan(&checkUser)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if checkUser != nil && checkUser.Uuid != getData.UUID {
		return berror.NewError(bcode.OperationFailed, "用户已存在")
	}
	// 编辑用户
	var getRole string
	if getData.IsAdmin {
		getRole = util.GetAdminRoleUUID()
	} else {
		getRole = util.GetUserRoleUUID()

	}
	_, err = dao.XfUser.Ctx(ctx).Data(do.XfUser{
		Username: getData.Username,
		Email:    getData.Email,
		Phone:    getData.Phone,
		Role:     getRole,
	}).Where(do.XfUser{Uuid: getData.UUID}).Update()
	// 如果不是管理员需要编辑公司
	if getData.IsAdmin {
		// 如果是管理员则删除公司
		_, err = dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: getData.UUID}).Delete()
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
	} else {
		// 检查公司是否存在
		var getCompany *entity.XfCompanies
		err = dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: getData.UUID}).Scan(&getCompany)
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		getModal := dao.XfCompanies.Ctx(ctx).Data(do.XfCompanies{
			Cods:           getData.CompanyCods,
			Name:           getData.Company,
			Address:        getData.CompanyAddress,
			Representative: getData.Representative,
			Uuid:           getData.UUID,
		})
		if getCompany == nil {
			_, err = getModal.Insert()
		} else {
			_, err = getModal.Where(do.XfCompanies{Cods: getData.CompanyCods}).Update()
		}
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
	}
	return nil
}

// UserGetList
//
// # 获取用户列表
//
// 获取用户列表接口, 用于获取用户列表. 用户在这里可以进行获取用户列表，获取用户列表的时候需要用户提供用户列表的信息，获取成功后会返回获取成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - userDetail: []*duser.UserDetailed, 用户详细信息
//   - err: error, 错误
func (s *sConsole) UserGetList(ctx context.Context) (userDetail []*duser.UserDetailed, err error) {
	g.Log().Noticef(ctx, "[LOGIC] 执行 UserGetList | 获取用户列表")
	// 获取用户列表
	var getUserList []*entity.XfUser
	err = dao.XfUser.Ctx(ctx).OrderDesc("created_at").Scan(&getUserList)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	getUserDetailedList := new([]*duser.UserDetailed)
	// 根据用户列表获取公司信息
	for _, user := range getUserList {
		ranCompany := new(dcompany.DCompany)
		if user.Role != util.GetAdminRoleUUID() {
			var getCompany *entity.XfCompanies
			err = dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: user.Uuid}).Scan(&getCompany)
			if err != nil {
				return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
			}
			if getCompany != nil {
				ranCompany = &dcompany.DCompany{
					Cods:           getCompany.Cods,
					Name:           getCompany.Name,
					Address:        getCompany.Address,
					Representative: getCompany.Representative,
				}
			}
		}
		// 根据角色获取角色名
		var getRole *entity.XfRole
		err = dao.XfRole.Ctx(ctx).Where(do.XfRole{Ruuid: user.Role}).Scan(&getRole)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		*getUserDetailedList = append(*getUserDetailedList, &duser.UserDetailed{
			User: duser.DUser{
				Uuid:     user.Uuid,
				Username: user.Username,
				Email:    user.Email,
				Phone:    user.Phone,
				Role:     getRole.Name,
			},
			Company: *ranCompany,
		})
	}
	return *getUserDetailedList, nil
}
