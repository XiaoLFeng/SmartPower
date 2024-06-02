package cmd

import (
	"SmartPower/internal/config/middleware"
	"SmartPower/internal/config/startup"
	"SmartPower/internal/config/task"
	"SmartPower/internal/controller/auth"
	"SmartPower/internal/controller/electricity"
	"SmartPower/internal/controller/user"
	"context"
	"github.com/bamboo-services/bamboo-utils/bmiddle"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 业务启动
			s := g.Server()
			// 初始化系统
			startup.SystemDatabaseTablePreparation(ctx)
			startup.SystemDataPreparation(ctx)
			startup.SystemConsolePreparation(ctx)
			startup.SystemPreparationFinal(ctx)
			// 启动定时器
			task.CleanTokenTask(ctx)
			// 路由表控制
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(bmiddle.BambooMiddleHandler)
				group.Middleware(middleware.MiddleAllowOrigin)
				group.Bind(
					auth.NewV1(),
					user.NewV1(),
					electricity.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
