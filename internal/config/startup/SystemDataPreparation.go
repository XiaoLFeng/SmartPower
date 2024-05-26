package startup

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

/*
SystemDataPreparation

# 初始化系统 - 系统数据准备

参数:
  - ctx: context.Context, 上下文
*/
func SystemDataPreparation(ctx context.Context) {
	// 准备信息表数据
	prepareInfoTableData(ctx, "system_name", "SmartPower")
	prepareInfoTableData(ctx, "system_version", "v1.0.0")
	prepareInfoTableData(ctx, "system_author", "xiao_lfeng")
	prepareInfoTableData(ctx, "system_description", "SmartPower 是一个智能电力管理系统")
	prepareInfoTableData(ctx, "system_github", "https://github.com/XiaoLFeng/SmartPower")
}

/*
prepareInfoTableData

# 准备信息表数据

参数:
  - ctx: context.Context, 上下文
  - key: string, 键
  - value: string, 值
*/
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
