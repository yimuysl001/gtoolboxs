package permission

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"os"
	"time"
)

func init() {
	logger.Logger.Trace("=======================设置有效期=========================")
	defer func() {
		logger.Logger.IfError(recover())
		logger.Logger.Trace("=======================设置有效期完成=========================")
	}()
	cron()
	getCtime()

}

func cron() {
	gcron.AddSingleton(gctx.New(), "@every 30m", func(ctx context.Context) {
		logger.Logger.Trace("permission check start")
		defer func() {
			logger.Logger.Trace("permission check stop")
			if r := recover(); r != nil {
				logger.Logger.ErrorCtx(ctx, r)
				//flags = false
				os.Exit(0)
			}

		}()
		now := time.Now()
		if ctime == "" {
			panic("程序未进行初始化，请先初始化:" + ckey)
		}
		if dtime.IsZero() {
			panic("程序初始化配置不正确:" + ckey)
		}
		if dtime.Before(now) {
			panic("程序已到期:" + dtime.String() + "，请联系管理员:" + ckey)
		}
	})

}
