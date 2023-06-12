package pongo

import (
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/flosch/pongo2/v6"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"strings"
	"time"
)

func init() {
	BuildFunction(sprig.FuncMap())

	BuildFunction(map[string]any{
		"uuid":          guid.S,       //唯一id 不需要参数 返回string
		"sql":           SqlMap,       //单条结果集返回 参数： [sql语句] [sql参数]... 返回map
		"sqlmust":       SqlMust,      //单条结果集返回(不可为空) 参数： [sql语句] [sql参数]... 返回map
		"sqls":          SqlMaps,      //多条结果集 参数： [sql语句] [sql参数]... 返回map切片
		"sqlsmust":      SqlsMust,     //多条结果集(不可为空) 参数： [sql语句] [sql参数]... 返回map切片
		"sqlTable":      sqlTable,     //单条结果集返回 参数： [sql语句] [sql参数]... 返回map
		"sqlTableMust":  sqlTableMust, //单条结果集返回(不可为空) 参数： [sql语句] [sql参数]... 返回map
		"sqlsTable":     sqlsTable,    //多条结果集 参数： [sql语句] [sql参数]... 返回map切片
		"sqlsTableMust": sqlsTableMust,
		"cache":         SqlCache, //字典数据缓存查询 参数：[表名] [条件] 条件参数  返回map
		"cat":           cat,      //字符拼接
		"datef":         date,     //时间格斯转换
		"todate": func(fmt, str string) time.Time {
			t, _ := time.ParseInLocation(fmt, str, time.Local)
			return t
		},

		"map":      rangMap,
		"sprintf":  fmt.Sprintf,
		"sendHttp": doHttp,
		"sendMq":   doMq,
		"insert":   InsertData,
		"delete":   DelData,
		"execsql":  Exec,
		"testShow": testShow,
		//"plugin":        plugin,
		"mapToJson": mapToJson,
		"mapToXml":  mapToXml,
		"strToMap":  strToMap,
		"PanicXX":   PanicXX,
		"toString":  gconv.String,
		"joinNo": func(s string, v interface{}) string {
			stringsl := strslice(v)
			var strs = make([]string, 0)
			for _, str := range stringsl {
				if len(str) > 0 {
					strs = append(strs, str)
				}
			}
			if len(strs) == 0 {
				return ""
			}
			return strings.Join(strs, s)
		},
	})
}

func BuildFunction(f map[string]any) {

	//for s, a := range f {
	//	funcmap[s] = a
	//}

	pongo2.DefaultSet.Globals.Update(f)
}

var funcmap = make(pongo2.Context)
