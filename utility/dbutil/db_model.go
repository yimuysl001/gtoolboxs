package dbutil

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"strings"
)

type dbModel struct {
	tablename string //表名称
	modelname string //使用模型
}

// 创建model
func New(model string) *dbModel {

	return &dbModel{
		tablename: model,
		modelname: model,
	}
}

// AS 数据库名称
func (d *dbModel) AS(name string) *dbModel {
	d.tablename = name
	return d
}

// 插入数据
func (d *dbModel) Insert(ctx context.Context, Tx gdb.TX, data interface{}) (sql.Result, error) {
	if d.modelname == d.tablename {
		return Tx.Insert(d.tablename, data)
	}
	//var sql string
	//var args []interface{}

	return Tx.Model(d.tablename).Fields(Tx.Model(d.modelname).GetFieldsStr()).Insert(data)

	//fields, err := Tx.Model().TableFields(d.modelname)
	//logger.Logger.PanicErrorCtx(ctx, err)
	//var sqlsb strings.Builder
	//var sqlvalue strings.Builder
	//
	//sql = "insert into " + d.tablename + "(%v) values(%v)"
	//sqlmap := gconv.Map(data)
	//for _, field := range fields {
	//	b, i := checkmapkey(field.Name, sqlmap)
	//	if !b {
	//		continue
	//	}
	//	sqlsb.WriteString(field.Name + ",")
	//	sqlvalue.WriteString("?,")
	//	args = append(args, i)
	//}
	//if len(args) < 1 {
	//	panic("未找到匹配字段")
	//}
	//sql = fmt.Sprintf(sql, sqlsb.String()[:sqlsb.Len()-1], sqlvalue.String()[:sqlvalue.Len()-1])
	//logger.Logger.TraceCtx(ctx, "sql:", sql)
	//
	//return Tx.Exec(sql, args...)

}

func checkmapkey(fieldname string, data map[string]interface{}) (bool, interface{}) {
	for k, v := range data {
		if strings.EqualFold(fieldname, k) {
			return true, v
		}
	}
	return false, nil

}