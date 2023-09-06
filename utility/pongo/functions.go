package pongo

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yimuysl001/gtoolboxs/utility/dbutil"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/mqutil"
	"github.com/yimuysl001/gtoolboxs/utility/myhttp"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var localdbname = ""

func SetDb(name string) {
	localdbname = name
}

func DelData(table string, where string, args ...interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()
	name, s, s2 := gettableName(table)
	_, err := dbutil.DB(name).Schema(s).Delete(context.Background(), s2, where, args...)

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
	s, s2 := getsqldb(sql)
	_, err := dbutil.DB(s).Exec(context.Background(), s2, args...)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	flag = true
	return
}

func ExecTran(sql string, args ...interface{}) (flag bool) {
	flag = false
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error(r)
			flag = false
		}
	}()
	s, s2 := getsqldb(sql)
	err := dbutil.Tran(gctx.New(), dbutil.DB(s), func(ctx2 context.Context, tx gdb.TX) {
		_, err := tx.Exec(s2, args...)
		logger.Logger.PanicErrorCtx(ctx2, err)
	})
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
	name, s, s2 := gettableName(table)
	err := dbutil.Tran(context.Background(), dbutil.DB(name).Schema(s), func(ctx2 context.Context, tx gdb.TX) {
		count := 0
		if key != "" {
			keydata := data[key]
			count, _ = tx.Model(s2).Count(key+"=?", keydata)
		}
		if count > 0 {
			keydata := data[key]
			delete(data, key)
			_, err := tx.Update(s2, data, key+"=?", keydata)
			//_, err := tx.Delete(s2, key+"=?", data[key])
			logger.Logger.PanicErrorCtx(ctx2, err)
		} else {
			_, err := tx.Insert(s2, data)
			logger.Logger.PanicErrorCtx(ctx2, err)
		}
	})

	logger.Logger.PanicErrorCtx(context.Background(), err)

	flag = true
	return
}

// SqlMap 查询单条数据
// sql 语句内容
// args 传参
func SqlMap(sql string, args ...interface{}) map[string]interface{} {
	s, s2 := getsqldb(sql)
	one, err := dbutil.DB(s).GetOne(context.Background(), s2, args...)
	if err != nil {
		panic(err)
	}

	return one.Map()

}

func sqlTable(tbname, where string, args ...interface{}) map[string]interface{} {
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).One(where, args)
	if err != nil {
		panic(err)
	}
	return one.Map()
}
func sqlTableMust(tbname, where string, args ...interface{}) map[string]interface{} {
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).One(where, args)
	if err != nil {
		panic(err)
	}
	if one.IsEmpty() {
		panic(fmt.Sprintf("未查询到数据，表=%v;条件=%v; 参数=%v", tbname, where, args))
	}
	return one.Map()
}

func sqlsTable(tbname, where string, args ...interface{}) []map[string]interface{} {
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).All(where, args)
	if err != nil {
		panic(err)
	}
	return one.List()
}
func sqlsTableMust(tbname, where string, args ...interface{}) []map[string]interface{} {
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).All(where, args)
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
	s, s2 := getsqldb(sql)
	one, err := dbutil.DB(s).GetOne(context.Background(), s2, args...)
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
	s, s2 := getsqldb(sql)
	all, err := dbutil.DB(s).GetAll(context.Background(), s2, args...)
	if err != nil {
		panic(err)
	}
	return all.List()

}

// SqlsMust 查询多条数据(不能为空)
// sql 语句内容
// args 传参
func SqlsMust(sql string, args ...interface{}) []map[string]interface{} {
	s, s2 := getsqldb(sql)
	all, err := dbutil.DB(s).GetAll(context.Background(), s2, args...)
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
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).Cache(gdb.CacheOption{Duration: time.Hour}).One(where, args)
	if err != nil {
		panic(err)
	}
	return one.Map()
}

// SqlsCache 缓存数据
// tbname 表名称
// where 查询条件
// args 传参
func SqlsCache(tbname string, where string, args ...interface{}) []map[string]interface{} {
	name, s, s2 := gettableName(tbname)
	one, err := dbutil.DB(name).Schema(s).Model(s2).Cache(gdb.CacheOption{Duration: time.Hour}).All(where, args)
	if err != nil {
		panic(err)
	}
	return one.List()
}

// gettableName
//
//	@Description:
//	@param tbname
//	@return string 配置名称
//	@return string  库名
//	@return string 表名
func gettableName(tbnamew string) (string, string, string) {
	db, tbname := getsqldb(tbnamew)

	if strings.Contains(tbname, "..") {
		ns := strings.Split(tbname, "..")
		if len(ns) > 2 {
			return ns[0], ns[1], ns[2]
		}
		n1 := strings.ToUpper(ns[1])

		ns2 := strings.SplitN(n1, ".DBO.", 2)
		if len(ns2) < 2 {
			return db, ns[0], ns[1]
		}
		return ns[0], ns2[0], ns2[1]

	}
	tbname = strings.ToUpper(tbname)
	ns2 := strings.SplitN(tbname, ".DBO.", 2)
	if len(ns2) < 2 {
		return db, "", tbname
	}

	return db, ns2[0], ns2[1]
}

func getsqldb(sqls string) (string, string) {
	regex := regexp.MustCompile(`\[\s*(\w+)\s*=\s*([^]]+)\s*]`)
	matches := regex.FindAllStringSubmatch(sqls, -1)
	var db = localdbname
	for _, match := range matches {
		// match[1] contains the key inside the square brackets (e.g., "DB")
		// match[2] contains the value inside the square brackets (e.g., "123")
		key := strings.TrimSpace(match[1])
		if strings.EqualFold(key, "DB") {
			db = strings.TrimSpace(match[2])
		}
	}

	sqls = regex.ReplaceAllString(sqls, "")
	return db, sqls

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

func rangMap2(k, v string, maps []map[string]interface{}) map[string]interface{} {
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
		m[gconv.String(key)] = value
	}

	return m

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

func mapToJson(data interface{}) string {

	return gjson.New(data).MustToJsonString()
}

func mapToXml(data interface{}, rootTag ...string) string {
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

func countInt(index int) []int {
	if index <= 0 {
		return nil
	}
	var sequence = make([]int, index+1)
	for i := 0; i < index; i++ {
		sequence[i] = i
	}
	return sequence

}

func customGenerateSequence(seqStart, seqEnd int) []int {
	var sequence []int
	for i := seqStart; i <= seqEnd; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}

func cfg(path string, name ...string) interface{} {
	return g.Cfg(name...).MustGet(context.Background(), path, nil).Val()
}
func addDate(d time.Time, i int, ds string) time.Time {

	switch strings.ToUpper(ds) {
	case "Y":
		d = d.AddDate(i, 0, 0)
	case "M":
		d = d.AddDate(0, i, 0)
	case "D":
		d = d.AddDate(0, 0, i)
	default:
		d = d.Add(time.Second * time.Duration(i))
	}

	return d
}

func GetType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

// list指定字段转为mao
func rangMapNo(k, v string, maps interface{}) map[string]interface{} {
	var m = make(map[string]interface{})

	if maps == nil {
		return m
	}
	switch va := maps.(type) {
	case []map[string]interface{}:
		return rangMap2(k, v, va)
	case []interface{}:
		for _, vs := range va {
			if vs == nil {
				continue
			}
			if m2, y := vs.(map[string]interface{}); y {
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
				m[gconv.String(key)] = value
			}
		}

	default:
		logger.Logger.PanicCtx(context.Background(), "未识别的数据类型")

	}

	//for _, m2 := range maps {
	//	key, ok := m2[k]
	//	if !ok {
	//		panic("未找到key字段：" + k)
	//	}
	//	if key == nil || key == "" {
	//		continue
	//	}
	//	value, ok := m2[v]
	//	if !ok {
	//		panic("未找到value字段：" + v)
	//	}
	//	m["K"+gconv.String(key)] = value
	//}

	return m

}

func rangMapg(k, v string, maps interface{}) map[string]interface{} {
	var m = make(map[string]interface{})

	if maps == nil {
		return m
	}
	switch va := maps.(type) {
	case []map[string]interface{}:
		return rangMap(k, v, va)
	case []interface{}:
		for _, vs := range va {
			if vs == nil {
				continue
			}
			if m2, y := vs.(map[string]interface{}); y {
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
		}

	default:
		logger.Logger.PanicCtx(context.Background(), "未识别的数据类型")

	}

	//for _, m2 := range maps {
	//	key, ok := m2[k]
	//	if !ok {
	//		panic("未找到key字段：" + k)
	//	}
	//	if key == nil || key == "" {
	//		continue
	//	}
	//	value, ok := m2[v]
	//	if !ok {
	//		panic("未找到value字段：" + v)
	//	}
	//	m["K"+gconv.String(key)] = value
	//}

	return m

}

func nowtime() (t time.Time) {
	defer func() {
		if err := recover(); err != nil || t.IsZero() {
			t = time.Now()
		}
	}()

	one, err := g.DB(localdbname).GetOne(context.Background(), "select getdate() as time")
	if err != nil {
		return time.Now()
	}

	if one.IsEmpty() {
		return time.Now()
	}

	return gconv.Time(one["time"], time.DateTime)

}
