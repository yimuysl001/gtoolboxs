package permission

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/sm4"
	"github.com/yimuysl001/gtoolboxs/utility/fileutil"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/timeutil"
	"os"
	"strings"
	"time"
)

func init() {
	//logger.Logger.Trace("=======================设置有效期=========================")
	//defer func() {
	//	logger.Logger.IfError(recover())
	//	logger.Logger.Trace("=======================设置有效期完成=========================")
	//}()
	//cron()
	//getCtime()
	logger.Logger.Trace("=======================设置有效期=========================")
	getCtime()
	cron()
	err := setCtime(ctime)
	if err == nil {
		logger.Logger.Info("有效期设置完成，有效期为：" + dtime.String())
		return
	}
	logger.Logger.Error(err)
	//Run()
	RunW()
	if !flags {
		logger.Logger.Trace("=======================设置失败=========================")
		os.Exit(0)
	}

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

func RunW() {
	for i := 0; i < 10; i++ {
		logger.Logger.Info("key:", ckey)
		//logger.Logger.Trace("key:", ckey)
		var t string
		fmt.Print("请输入注册码：")
		_, err := fmt.Scan(&t)
		if err != nil {
			logger.Logger.Error("输入错误:", err)
			logger.Logger.Error("再错 ", 9-i, "次后程序关闭")
			continue
		}
		err = setCtime(t)
		if err != nil {
			logger.Logger.Error("设置出错：", err)
			logger.Logger.Error("再错 ", 9-i, "次后程序关闭")
			continue
		}
		flags = true

		logger.Logger.Info("有效期设置完成，有效期为：" + dtime.String())
		break
	}

}

// 设置权限日期
func setCtime(str string) (errn error) {
	str = strings.TrimSpace(str)

	if str == "请输入注册码" || str == "" {
		return errors.New("未设置注册码")
	}
	if str == ckey {
		dtime = time.Now().AddDate(0, 0, 1)
		logger.Logger.Trace("临时权限")
		return nil
	}
	defer func() {
		if err := recover(); err != nil {
			logger.Logger.Error(err)
			errn = errors.New(fmt.Sprintf("%v", err))
		}
	}()

	base64 := sm4.DectyptEcbBase64(str)
	//key 解密
	ecbBase64 := sm4.DectyptByKeyEcbBase64(string(base64), ckey)
	dtime = timeutil.StrToTime(string(ecbBase64))
	if dtime.IsZero() {
		return errors.New("注册码输入不正确！！！")
	}

	if dtime.Before(time.Now()) {
		return errors.New("设置有效期:" + dtime.String() + "，在当前时间之前。")
	}
	exists, _ := fileutil.PathExists(configpath)
	if !exists {
		os.MkdirAll(configpath, 0666)
	}

	return os.WriteFile(permissionpath, []byte(str), 0666)

}
