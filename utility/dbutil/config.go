package dbutil

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

// TBDBCONFIG 数据库配置
type TBDBCONFIG struct {
	CYWBM  string `json:"cywbm"`  // 配置业务编码
	STATUS int    `json:"status"` // 状态 1 使用
	Host   string `json:"host"`   // Host of server, ip or domain like: 127.0.0.1, localhost
	Port   string `json:"port"`   // Port, it's commonly 3306.
	User   string `json:"user"`   // Authentication username.
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
		logger.Logger.PanicCtx(ctx, name+"数据库配置为设置")
	}
	var tb = make([]TBDBCONFIG, 0)

	err = all.Structs(&tb)
	logger.Logger.PanicErrorCtx(ctx, err)

	var group = make(gdb.ConfigGroup, len(tb))
	for i, tbdbconfig := range tb {
		var newcf = g.DB().GetConfig()
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
		newcf.User = tbdbconfig.User
		//附加属性
		//newcf.Extra = "app name=" + name + "测试"
		newcf.Extra = tbdbconfig.Extra
		if tbdbconfig.Role == "" {
			tbdbconfig.Role = "master"
		}
		newcf.Role = tbdbconfig.Role
		group[i] = *newcf
	}

	return group

}
