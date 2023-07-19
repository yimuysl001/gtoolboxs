package dbutil

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yimuysl001/gtoolboxs/utility/cache"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"sync"
)

const (
	cachekey = "db.name.dbcache."
	gfcache  = "gf.core.component.database"
)

var once sync.Once

// DB db获取
func DB(name ...string) (d gdb.DB) {
	//如果 name 为空，直接获取本地配置
	if name == nil || len(name) < 1 || name[0] == "" {
		return g.DB()
	}
	defer func() {
		if d != nil && cache.GetAdapter() != nil {
			once.Do(func() {
				d.GetCache().SetAdapter(cache.GetAdapter())
			})
		}
	}()

	n := name[0]

	//尝试获取本地配置
	mdb, ok := dbname(n)
	if ok { //获取成功，直接返回本地数据库配置
		return mdb
	}
	//设置数据库配置
	SetDb(n)

	return db(n)

}

func dbname(name ...string) (gd gdb.DB, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			ok = false
		}
	}()
	return db(name...), true

}

// SetDb 设置数据源
func SetDb(name string) {
	if setLocalDb(name) {
		return
	}
	//config := db().GetConfig()
	//newcf := *config
	////todo 配置处理
	////newcf.Link = "sqlite: " + name
	////数据库ip
	//newcf.Host = "192.168.200.26"
	////数据库种类
	//newcf.Type = "mssql"
	////数据库密码
	//newcf.Pass = "123qwe,."
	////端口
	//newcf.Port = "1433"
	////选择库
	//newcf.Name = "YXHIS"
	////数据库登录名
	//newcf.User = "sa"
	////附加属性
	//newcf.Extra = "app name=" + name + "测试"
	//GetLink(&newcf)
	gdb.SetConfigGroup(name, GetConfigGroup(gctx.New(), name))
	if cache.GetAdapter() != nil {
		instance, err := gdb.Instance(name)
		logger.Logger.PanicErrorCtx(context.Background(), err)
		instance.GetCache().SetAdapter(cache.GetAdapter())
	}

}

func setLocalDb(name string) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			ok = false
		}
	}()
	_ = g.DB(name)
	config := gdb.GetConfig(name)
	if config == nil || len(config) < 1 {
		return false
	}
	//config := g.DB(name).GetConfig()
	//logger.Logger.Info(config)
	gdb.SetConfigGroup(name, config)
	if cache.GetAdapter() != nil {
		instance, err := gdb.Instance(name)
		logger.Logger.PanicErrorCtx(context.Background(), err)
		instance.GetCache().SetAdapter(cache.GetAdapter())
	}
	return true
}

func db(name ...string) gdb.DB {
	instance, err := gdb.Instance(name...)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	return instance
}
