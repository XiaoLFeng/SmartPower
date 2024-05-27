package startup

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"context"
	"crypto/md5"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// SystemDataPreparation
//
// # 初始化系统
//
// 初始化系统数据, 用于初始化系统数据. 系统数据初始化包括系统信息表数据, 角色表数据等.
//
// # 参数:
//   - ctx: context.Context, 上下文
func SystemDataPreparation(ctx context.Context) {
	glog.Noticef(ctx, "[STARTUP] 开始进行系统数据准备检查")
	// 准备信息表数据
	prepareInfoTableData(ctx, "system_name", "SmartPower")
	prepareInfoTableData(ctx, "system_version", "v1.0.0")
	prepareInfoTableData(ctx, "system_author", "xiao_lfeng")
	prepareInfoTableData(ctx, "system_description", "一个电力管理系统")
	prepareInfoTableData(ctx, "system_github", "https://github.com/XiaoLFeng/SmartPower")

	prepareRoleTableData(ctx, md5.Sum([]byte("admin")), "admin", "超级管理员", "拥有系统所有权限")
	prepareRoleTableData(ctx, md5.Sum([]byte("duser")), "duser", "普通用户", "普通用户权限")
}

// prepareInfoTableData
//
// # 准备信息表数据
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - key: string, 键
//   - value: string, 值
func prepareInfoTableData(ctx context.Context, key string, value string) {
	result, _ := dao.XfInfo.Ctx(ctx).Where(do.XfInfo{Keyword: key}).One()
	if result.IsEmpty() {
		glog.Debugf(ctx, "[STARTUP] 准备信息表数据: [%s]%s", key, value)
		newUUID, _ := uuid.NewV7()
		_, err := dao.XfInfo.Ctx(ctx).
			Data(do.XfInfo{
				SystemId:  newUUID.String(),
				Keyword:   key,
				Value:     value,
				UpdatedAt: gtime.Now(),
			}).Insert()
		if err != nil {
			glog.Warningf(ctx, "[STARTUP] 准备信息表数据失败: %s", err.Error())
		}
	}
}

// prepareRoleTableData
//
// # 准备角色表数据
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - ruuid: uuid.UUID, 角色UUID
//   - rname: string, 角色名称
//   - displayName: string, 显示名称
//   - desc: string, 描述
func prepareRoleTableData(ctx context.Context, ruuid uuid.UUID, rname string, displayName string, desc string) {
	var getRole *entity.XfRole
	err := dao.XfRole.Ctx(ctx).Where(do.XfRole{Ruuid: ruuid.String()}).Scan(&getRole)
	if err == nil {
		if getRole == nil {
			_, err := dao.XfRole.Ctx(ctx).Data(do.XfRole{
				Ruuid:       ruuid.String(),
				Name:        rname,
				DisplayName: displayName,
				Description: desc,
			}).Insert()
			if err != nil {
				glog.Warningf(ctx, "[STARTUP] 准备角色表数据失败: %s", err.Error())
			}
		}
	}
}
