package task

import (
	"SmartPower/internal/dao"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

// CleanTokenTask
//
// # 清理过期 Token
//
// 清理过期 Token 任务, 用于定时清理过期的 Token. 任务执行时间为每5分钟执行一次.
//
// # 参数:
//   - ctx: context.Context, 上下文
func CleanTokenTask(ctx context.Context) {
	gtimer.AddTimes(ctx, time.Minute, 5, func(_ context.Context) {
		startTime := gtime.Now().TimestampMilli()
		_, _ = dao.XfToken.Ctx(ctx).Where("expire_time < ?", gtime.Now()).Delete()
		glog.Noticef(
			ctx,
			"[TASK] 执行 CleanTokenTask | 清理过期 Token 完成 | 耗时: %d ms",
			gtime.Now().TimestampMilli()-startTime,
		)
	})
}
