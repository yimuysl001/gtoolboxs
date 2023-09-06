package dbutil

import (
	"context"
	"database/sql"
	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"  //加载数据数据库驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"  //加载数据数据库驱动
	_ "github.com/gogf/gf/contrib/drivers/oracle/v2" //加载数据数据库驱动
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"  //加载数据数据库驱动
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2" //加载数据数据库驱动
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"regexp"
	"strings"
)

// mssql: user id=账号;password=密码;server=地址;port=端口;database=数据库名称;encrypt=disable
const mssqllink = "user id=%v;password=%v;server=%v;port=%v;database=%v;encrypt=disable"

// sqlite: 文件绝对路径 (如: /var/lib/db.sqlite3)
const sqlitelink = "%v"

// oracle 账号/密码@地址:端口/数据库名称
const oraclelink = "%v/%v@%v:%v/%v"

// pgsql: user=账号 password=密码 host=地址 port=端口 dbname=数据库名称
const pgsqllink = "user=%v password=%v host=%v port=%v dbname=%v"

// mysql: 账号:密码@tcp(地址:端口)/数据库名称
const mysqllink = "%v:%v@tcp(%v:%v)/%v"

func init() {

	err := g.DB().PingSlave()
	if err == nil {
		logger.Logger.Info("本地配置数据库连接成功")
	} else {
		logger.Logger.Error("本地配置数据库连接失败：", err)
	}

}
func GetHookHandler(group string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			//logger.Logger.InfoCtx(ctx, "查询语句：", in.Sql)
			//logger.Logger.InfoCtx(ctx, "查询表名：", in.Table)
			//in.Model.Cache(gdb.CacheOption{Duration: time.Hour, Force: true})
			if DB(group).GetConfig().Type == "mssql" {
				in.Sql = strings.ReplaceAll(in.Sql, in.Table, in.Table+" with(nolock) ")
				re := regexp.MustCompile(`"([^"]+)"`)
				in.Sql = re.ReplaceAllString(in.Sql, `"[$1]"`)

			}

			result, err = in.Next(ctx)
			//logger.Logger.InfoCtx(ctx, "==========查询结果===========")
			return
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			in.Model.Master()
			result, err = in.Next(ctx)

			if err == nil {
				logger.Logger.IfErrorCtx(ctx, DB(group).GetCore().ClearCache(ctx, in.Table))

			}
			return
		},

		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			in.Model.Master()
			result, err = in.Next(ctx)
			if err == nil {
				logger.Logger.IfErrorCtx(ctx, DB(group).GetCore().ClearCache(ctx, in.Table))
			}
			return
		},

		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			in.Model.Master()
			result, err = in.Next(ctx)
			if err == nil {
				logger.Logger.IfErrorCtx(ctx, DB(group).GetCore().ClearCache(ctx, in.Table))
			}
			return
		},
	}

}
