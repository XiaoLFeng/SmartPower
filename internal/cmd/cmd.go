package cmd

import (
	"SmartPower/internal/config/middleware"
	"SmartPower/internal/config/startup"
	"SmartPower/internal/controller/auth"
	"context"

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
			// 路由表控制
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.HandlerResponseMiddleware)
				group.Bind(
					auth.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
