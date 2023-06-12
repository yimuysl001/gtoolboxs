package pongo

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yimuysl001/gtoolboxs/utility/dbutil"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/mqutil"
	"github.com/yimuysl001/gtoolboxs/utility/myhttp"

	"reflect"
	"strings"
	"time"
)

func DelData(table string, where string, args ...interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()

	_, err := dbutil.DB().Delete(context.Background(), table, where, args...)

	logger.Logger.PanicErrorCtx(context.Background(), err)

	flag = true
	return
}

func Exec(sql string, args ...interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()
	_, err := dbutil.DB().Exec(context.Background(), sql, args...)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	flag = true
	return
}

func InsertData(table string, key string, data map[string]interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()

	err := dbutil.Tran(context.Background(), dbutil.DB(), func(ctx2 context.Context, tx gdb.TX) {
		if key != "" {
			_, err := tx.Delete(table, key+"=?", data[key])
			logger.Logger.PanicErrorCtx(ctx2, err)
		}
		_, err := tx.Insert(table, data)
		logger.Logger.PanicErrorCtx(ctx2, err)

	})

	logger.Logger.PanicErrorCtx(context.Background(), err)

	flag = true
	return
}

// SqlMap 查询单条数据
// sql 语句内容
// args 传参
func SqlMap(sql string, args ...interface{}) map[string]interface{} {
	one, err := dbutil.DB().GetOne(context.Background(), sql, args...)
	if err != nil {
		panic(err)
	}

	return one.Map()

}

func sqlTable(tbname, where string, args ...interface{}) map[string]interface{} {
	one, err := dbutil.DB().Model(tbname).One(where, args)
	if err != nil {
		panic(err)
	}
	return one.Map()
}
func sqlTableMust(tbname, where string, args ...interface{}) map[string]interface{} {
	one, err := dbutil.DB().Model(tbname).One(where, args)
	if err != nil {
		panic(err)
	}
	if one.IsEmpty() {
		panic(fmt.Sprintf("未查询到数据，表=%v;条件=%v; 参数=%v", tbname, where, args))
	}
	return one.Map()
}

func sqlsTable(tbname, where string, args ...interface{}) []map[string]interface{} {
	one, err := dbutil.DB().Model(tbname).All(where, args)
	if err != nil {
		panic(err)
	}
	return one.List()
}
func sqlsTableMust(tbname, where string, args ...interface{}) []map[string]interface{} {
	one, err := dbutil.DB().Model(tbname).All(where, args)
	if err != nil {
		panic(err)
	}
	if one.IsEmpty() {
		panic(fmt.Sprintf("未查询到数据，表=%v;条件=%v; 参数=%v", tbname, where, args))
	}
	return one.List()
}

// SqlMust 查询单条数据(不能为空)
// sql 语句内容
// args 传参
func SqlMust(sql string, args ...interface{}) map[string]interface{} {
	one, err := dbutil.DB().GetOne(context.Background(), sql, args...)
	if err != nil {
		panic(err)
	}
	if one.IsEmpty() {
		panic(fmt.Sprintf("未查询到数据，sql=%v;参数%v", sql, args))
	}
	return one.Map()

}

// SqlMaps 查询多条数据
// sql 语句内容
// args 传参
func SqlMaps(sql string, args ...interface{}) []map[string]interface{} {
	all, err := dbutil.DB().GetAll(context.Background(), sql, args...)
	if err != nil {
		panic(err)
	}
	return all.List()

}

// SqlsMust 查询多条数据(不能为空)
// sql 语句内容
// args 传参
func SqlsMust(sql string, args ...interface{}) []map[string]interface{} {
	all, err := dbutil.DB().GetAll(context.Background(), sql, args...)
	if err != nil {
		panic(err)
	}
	if all.IsEmpty() {
		panic(fmt.Sprintf("未查询到数据，sql=%v;参数%v", sql, args))
	}

	return all.List()

}

// SqlCache 缓存数据
// tbname 表名称
// where 查询条件
// args 传参
func SqlCache(tbname string, where string, args ...interface{}) map[string]interface{} {
	one, err := dbutil.DB().Model(tbname).Cache(gdb.CacheOption{Duration: time.Hour}).One(where, args)
	if err != nil {
		panic(err)
	}
	return one.Map()
}

// ProSql 存储过程执行
//funcs ProSql(sql string, args ...interface{}) map[string]interface{} {
//	one, err := dbutil.DB().GetOne(context.Background(), sql, args...)
//	if err != nil {
//		panic(err)
//	}
//	if one.IsEmpty() {
//		panic(fmt.Sprintf("存储过程未查询到数据，sql=%v;参数%v", sql, args))
//	}
//
//	return one.Map()
//
//}

func cat(v ...interface{}) string {
	v = removeNilElements(v)
	r := strings.TrimSpace(strings.Repeat("%v", len(v)))
	return fmt.Sprintf(r, v...)
}

func removeNilElements(v []interface{}) []interface{} {
	newSlice := make([]interface{}, 0, len(v))
	for _, i := range v {
		if i != nil {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

func date(fmt string, date interface{}) string {

	if date == nil || date == "" {
		return ""
	}
	//logger.Logger.Info(fmt)
	//logger.Logger.Info(date)
	zone := dateInZone(fmt, date, "Local")
	//logger.Logger.Info(zone)
	return zone
}

func rangMap(k, v string, maps []map[string]interface{}) map[string]interface{} {
	var m = make(map[string]interface{})
	for _, m2 := range maps {
		key, ok := m2[k]
		if !ok {
			panic("未找到key字段：" + k)
		}
		if key == nil || key == "" {
			continue
		}
		value, ok := m2[v]
		if !ok {
			panic("未找到value字段：" + v)
		}
		m["K"+gconv.String(key)] = value
	}

	return m

}

func dateInZone(fmt string, date interface{}, zone string) string {
	var t time.Time
	switch d := date.(type) {
	default:
		t = gconv.Time(d)
	case time.Time:
		t = d
	case *time.Time:
		t = *d
	case int64:
		t = time.Unix(d, 0)
	case int:
		t = time.Unix(int64(d), 0)
	case int32:
		t = time.Unix(int64(d), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return t.In(loc).Format(fmt)
}

func doHttp(url string, data interface{}) (bake map[string]interface{}) {
	bake = make(map[string]interface{})

	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			bake["flag"] = false
		}
	}()
	post := myhttp.HttpPost(context.Background(), url, nil, data)
	logger.Logger.Trace("返回数据:", post)

	bake = gjson.New(post).Map()
	bake["flag"] = true
	return
}

func doMq(cywbm string, data interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()
	err := mqutil.MQ(cywbm).Publish(gconv.Bytes(data))
	logger.Logger.PanicErrorCtx(context.Background(), err)

	flag = true
	return
}

func testShow(str interface{}) interface{} {

	logger.Logger.Info(str)

	return str
}

func mapToJson(data map[string]interface{}) string {

	return gjson.New(data).MustToJsonString()
}

func mapToXml(data map[string]interface{}, rootTag ...string) string {
	return gjson.New(data).MustToXmlString(rootTag...)
}

func strToMap(data interface{}) map[string]interface{} {
	return gjson.New(data).Map()
}

func PanicXX(i interface{}) interface{} {
	panic(i)
	return i
}

func strslice(v interface{}) []string {
	switch v := v.(type) {
	case []string:
		return v
	case []interface{}:
		b := make([]string, 0, len(v))
		for _, s := range v {
			if s != nil {
				b = append(b, gconv.String(s))
			}
		}
		return b
	default:
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Array, reflect.Slice:
			l := val.Len()
			b := make([]string, 0, l)
			for i := 0; i < l; i++ {
				value := val.Index(i).Interface()
				if value != nil {
					b = append(b, gconv.String(value))
				}
			}
			return b
		default:
			if v == nil {
				return []string{}
			}

			return []string{gconv.String(v)}
		}
	}
}
