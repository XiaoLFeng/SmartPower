package cmd

import (
	"SmartPower/internal/config/middleware"
	"SmartPower/internal/config/startup"
	"SmartPower/internal/config/task"
	"SmartPower/internal/controller/auth"
	"SmartPower/internal/controller/console"
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
			// 前端路由表
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.GET("/", func(r *ghttp.Request) { r.Response.ServeFile("resource/public/index.html") })
				group.GET("/auth/*", func(r *ghttp.Request) { r.Response.ServeFile("resource/public/index.html") })
				group.GET("/home/*", func(r *ghttp.Request) { r.Response.ServeFile("resource/public/index.html") })
				group.GET("/console/*", func(r *ghttp.Request) { r.Response.ServeFile("resource/public/index.html") })
			})
			// 后端路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(bmiddle.BambooMiddleHandler)
				group.Middleware(middleware.MiddleAllowOrigin)
				group.Bind(
					auth.NewV1(),
					user.NewV1(),
					electricity.NewV1(),
					console.NewV1(),
				)
			})
			s.AddStaticPath("/assets", "resource/public/assets")
			s.AddStaticPath("/sql", "resource/sql")
			s.Run()
			return nil
		},
	}
)
