package startup

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"context"
	"encoding/base64"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

/*
SystemConsolePreparation

# 初始化系统 - 超级管理员准备

参数:
  - ctx: context.Context, 上下文
*/
func SystemConsolePreparation(ctx context.Context) {
	glog.Debugf(ctx, "[STARTUP] 开始进行超级管理员准备检查")
	userUuid, _ := uuid.FromBytes([]byte("SmartPowerSuperConsoleAdminUuid"))
	var getUser *do.XfUser
	_ = dao.XfUser.Ctx(ctx).Where(do.XfUser{Uuid: userUuid.String()}).Scan(&getUser)
	if getUser == nil {
		glog.Noticef(ctx, "[STARTUP] 超级管理员不存在, 创建超级管理员")
		glog.Noticef(ctx, "\t> 账号：admin")
		glog.Noticef(ctx, "\t> 密码：admin")
		buildPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(base64.StdEncoding.EncodeToString([]byte("admin"))),
			bcrypt.DefaultCost,
		)
		_, err := dao.XfUser.Ctx(ctx).Data(do.XfUser{
			Uuid:     userUuid.String(),
			Username: "admin",
			Email:    "admin@sp.com",
			Phone:    "18888888888",
			Password: string(buildPassword),
		}).Insert()
		if err != nil {
			glog.Errorf(ctx, "[STARTUP] 创建超级管理员失败: %s", err.Error())
		}
	}
}
