package pongo

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/beevik/etree"
	"github.com/flosch/pongo2/v6"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/yimuysl001/gtoolboxs/utility/cipher/aesutil"
	"github.com/yimuysl001/gtoolboxs/utility/xmlutil"
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
		"cache":         SqlCache,  //字典数据缓存查询 参数：[表名] [条件] 条件参数  返回map
		"caches":        SqlsCache, //字典数据缓存查询 参数：[表名] [条件] 条件参数  返回map
		"cat":           cat,       //字符拼接
		"datef":         date,      //时间格斯转换
		"todate": func(fmt, str string) time.Time {
			t, _ := time.ParseInLocation(fmt, str, time.Local)
			return t
		},
		"addDate":   addDate,
		"map":       rangMapg,
		"listToMap": rangMapNo,
		"getType":   GetType,
		"sprintf":   fmt.Sprintf,
		"sendHttp":  doHttp,
		"sendMq":    doMq,
		"insert":    InsertData,
		"delete":    DelData,
		"execsql":   Exec,
		"exectran":  ExecTran,
		"testShow":  testShow,
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
		"now": nowtime,
		"formatData": func(data string) string {
			return IndexDataMust(gctx.New(), data)
		},
		"replaceN": strings.Replace,
		"getCfg":   cfg,
		"gmd5":     gmd5.MustEncrypt,
		"aesEcb": func(key string, data interface{}) string {
			encrypted := aesutil.AesEncryptECB(gconv.Bytes(data), gconv.Bytes(key))
			return hex.EncodeToString(encrypted)
		},
		"substr": func(ks int, js int, body interface{}) string {
			s := gconv.String(body)
			if s == "" {
				return s
			}
			return s[ks:js]

		},
		"GetTimestamp": func() string {
			return gconv.String(time.Now().UnixMilli())
		},
		"getRoot": func(body string) *etree.Element {
			return xmlutil.GetROOT(body)

		},
		"getXpathData": func(root *etree.Element, path string) string {
			return xmlutil.GetElementValue(root, path)
		},
		"getElement": func(root *etree.Element, path string) *etree.Element {
			return xmlutil.GetElement(root, path)
		},
		"getElements": func(root *etree.Element, path string) []*etree.Element {
			return xmlutil.GetElements(root, path)
		},
	})
	pongo2.RegisterFilter("error", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		return nil, &pongo2.Error{
			Sender:    "error",
			OrigError: errors.New(in.String()),
		}
	})

}

func BuildFunction(f map[string]any) {

	//for s, a := range f {
	//	funcmap[s] = a
	//}

	pongo2.DefaultSet.Globals.Update(f)
}

func GetFuncs() map[string]any {
	return pongo2.DefaultSet.Globals
}
