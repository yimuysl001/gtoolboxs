package dbutil

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"regexp"
	"strings"
)

// TBDBCONFIG 数据库配置
type TBDBCONFIG struct {
	CYWBM  string `json:"cywbm"`  // 配置业务编码
	STATUS int    `json:"status"` // 状态 1 使用
	Host   string `json:"host"`   // Host of server, ip or domain like: 127.0.0.1, localhost
	Port   string `json:"port"`   // Port, it's commonly 3306.
	CUser  string `json:"cuser"`  // Authentication username.
	Pass   string `json:"pass"`   // Authentication password.
	Name   string `json:"name"`   // Default used database name.
	Type   string `json:"type"`   // Database type: mysql, sqlite, mssql, pgsql, oracle.
	Link   string `json:"link"`   // (Optional) Custom link information for all configuration in one single string.
	Extra  string `json:"extra"`  // (Optional) Extra configuration according the registered third-party database driver.
	Role   string `json:"role"`   // (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
}

func GetConfigGroup(ctx context.Context, name string) gdb.ConfigGroup {
	all, err := g.DB().Model("TBDBCONFIG").All("status=1 and CYWBM=?", name)
	logger.Logger.PanicErrorCtx(ctx, err)
	if all.IsEmpty() {
		logger.Logger.PanicCtx(ctx, name+"数据库配置未设置")
	}
	var tb = make([]TBDBCONFIG, 0)

	err = all.Structs(&tb)
	logger.Logger.PanicErrorCtx(ctx, err)

	var group = make(gdb.ConfigGroup, len(tb))
	var conf = g.DB().GetConfig()
	for i, tbdbconfig := range tb {
		var newcf = *conf
		newcf.Link = tbdbconfig.Link
		//数据库ip
		newcf.Host = tbdbconfig.Host
		//数据库种类
		newcf.Type = tbdbconfig.Type
		//数据库密码
		newcf.Pass = tbdbconfig.Pass
		//端口
		newcf.Port = tbdbconfig.Port
		//选择库
		newcf.Name = tbdbconfig.Name
		//数据库登录名
		newcf.User = tbdbconfig.CUser
		//附加属性
		//newcf.Extra = "app name=" + name + "测试"
		//if tbdbconfig.Extra != "" {
		newcf.Extra = tbdbconfig.Extra
		//}

		debug, b := setDebug(tbdbconfig.Extra)
		if debug {
			newcf.Debug = b
			newcf.Extra = reDebug(tbdbconfig.Extra)
		}

		if tbdbconfig.Role == "" {
			tbdbconfig.Role = "master"
		}

		newcf.Role = tbdbconfig.Role
		group[i] = newcf
	}

	return group

}

func reDebug(extra string) string {
	re := regexp.MustCompile(`\bdebug=[^;]*;?`)
	result1 := re.ReplaceAllString(extra, "")
	return strings.TrimRight(result1, ";")
}

func setDebug(extra string) (bool, bool) {
	extra = strings.TrimSpace(extra)
	if extra == "" {
		return false, false
	}
	parts := strings.Split(extra, ";")
	for _, part := range parts {
		if strings.Contains(part, "=") {
			keyValue := strings.SplitN(part, "=", 2)
			key := keyValue[0]
			value := keyValue[1]
			if !strings.EqualFold("debug", key) {
				continue
			}

			if strings.EqualFold("true", value) {
				return true, true
			} else {
				return true, false
			}

		}
	}

	return false, false
}
