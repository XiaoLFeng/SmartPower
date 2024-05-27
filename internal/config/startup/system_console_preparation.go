package startup

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/util"
	"context"
	"crypto/md5"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
)

/*
SystemConsolePreparation

# 初始化系统 - 超级管理员准备

参数:
  - ctx: context.Context, 上下文
*/
func SystemConsolePreparation(ctx context.Context) {
	glog.Debugf(ctx, "[STARTUP] 开始进行超级管理员准备检查")
	var getUser *do.XfUser
	_ = dao.XfUser.Ctx(ctx).
		Where(do.XfUser{Uuid: uuid.UUID(md5.Sum([]byte("SmartPowerSuperConsoleUser"))).String()}).
		Scan(&getUser)
	if getUser == nil {
		glog.Noticef(ctx, "[STARTUP] 超级管理员不存在, 创建超级管理员")
		glog.Noticef(ctx, "\t> 账号：admin")
		glog.Noticef(ctx, "\t> 密码：admin")
		_, err := dao.XfUser.Ctx(ctx).Data(do.XfUser{
			Uuid:     uuid.UUID(md5.Sum([]byte("SmartPowerSuperConsoleUser"))).String(),
			Username: "admin",
			Email:    "admin@sp.com",
			Phone:    "18888888888",
			Role:     uuid.UUID(md5.Sum([]byte("admin"))).String(),
			Password: util.EncodePassword("admin"),
		}).Insert()
		if err != nil {
			glog.Errorf(ctx, "[STARTUP] 创建超级管理员失败: %s", err.Error())
		}
	}
}
