package startup

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"strings"
)

/*
SystemDatabaseTablePreparation

# 初始化系统 - 系统数据库表准备

参数:
  - ctx: context.Context, 上下文
*/
func SystemDatabaseTablePreparation(ctx context.Context) {
	// 数据库检查准备
	prepareDatabase(ctx, "xf_info")
	prepareDatabase(ctx, "xf_user")
	prepareDatabase(ctx, "xf_companies")
	prepareDatabase(ctx, "xf_electricity_rates")
}

/*
prepareDatabase

# 准备数据库

参数:
  - tableName: string, 表名
*/
func prepareDatabase(ctx context.Context, tableName string) {
	glog.Debugf(ctx, "[STARTUP] 检查数据表: %s", tableName)
	result, _ := g.DB().GetOne(ctx, "SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = ?", tableName)
	if result == nil {
		glog.Noticef(ctx, "[STARTUP] 数据表 %s 缺失，正在创建", tableName)
		// 读取文件
		getFileString := gfile.GetContents("resource/sql/" + tableName + ".sql")
		getContent := strings.Split(getFileString, "go")
		// 数据表不存在创建数据表
		for _, content := range getContent {
			_, err := g.DB().Exec(ctx, content)
			if err != nil {
				g.Throw(err.Error())
			}
		}
	}
}
