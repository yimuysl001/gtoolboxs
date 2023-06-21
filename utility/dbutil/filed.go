package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/yimuysl001/gtoolboxs/utility/logger"

	"net/url"
)

// GetSqlFiled 获取数据字段
func GetSqlFiled(name string, ASQL string, args ...interface{}) ([]string, error) {
	logger.Logger.Trace(ASQL)
	//db, err := GetSyncDB(name)

	db, err := GetSyncDB(name)
	if err != nil {
		return nil, err
	}
	query, err := db.Query(ASQL, args...)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	columns, err := query.ColumnTypes()
	if err != nil {
		return nil, err
	}
	var names = make([]string, len(columns))

	for i, column := range columns {
		names[i] = column.Name()
	}

	return names, nil

}

func GetSyncDB(name string) (*sql.DB, error) {

	//cdb := sdb.GetCache(name)
	//if cdb != nil {
	//	s, ok := cdb.(*sql.DB)
	//	if !ok {
	//		return nil, gerror.New("未获取数据库相应数据")
	//	}
	//	return s, nil
	//}
	Type, Link := getnewconf(name)

	open, err := sql.Open(Type, Link)
	if err != nil {
		return nil, err
	}
	err = open.Ping()
	if err != nil {
		return nil, err
	}
	//sdb.SetCache(name, open)

	return open, nil

}

// 获取数据库相应
func getnewconf(name string) (string, string) {

	//consts.YXHISDB

	config := DB(name).GetConfig()
	if config == nil {
		panic("未找到" + name + "数据库配置")
	}
	if config.Type == "" {
		panic(name + "数据库配置有错误")
	}
	link := config.Link
	if link != "" {
		return config.Type, link
	}

	return config.Type, GetLink(config) //fmt.Sprintf(mssqllink, config.yaml.User, config.yaml.Pass, config.yaml.Host, config.yaml.Port, config.yaml.Name, consts.AppName)
}

func GetLink(node *gdb.ConfigNode) string {

	switch node.Type {
	case "mysql":
		if node.Port == "" {
			node.Port = "3306"
		}
		node.Link = fmt.Sprintf(mysqllink, node.User, node.Pass, node.Host, node.Port, node.Name)
		if node.Charset != "" {
			node.Link = fmt.Sprintf("%s?charset=%s", node.Link, url.QueryEscape(node.Charset))
		}
		if node.Timezone != "" {
			node.Link = fmt.Sprintf("%s&loc=%s", node.Link, url.QueryEscape(node.Timezone))
		}
		if node.Extra != "" {
			node.Link = fmt.Sprintf("%s&%s", node.Link, node.Extra)
		}
	case "pgsql":
		if node.Port == "" {
			node.Port = "5432"
		}
		node.Link = fmt.Sprintf(pgsqllink, node.User, node.Pass, node.Host, node.Port, node.Name)

	case "sqlite":
		if node.Link == "" {
			node.Link = fmt.Sprintf(sqlitelink, node.Host)
		}

	case "oracle":
		if node.Port == "" {
			node.Port = "1521"
		}
		node.Link = fmt.Sprintf(oraclelink, node.User, node.Pass, node.Host, node.Port, node.Name)
	default: //默认mssql
		if node.Port == "" {
			node.Port = "1433"
		}
		if node.Extra == "" {
			node.Extra = "app name=获取表字段"
		}

		node.Link = fmt.Sprintf(mssqllink, node.User, node.Pass, node.Host, node.Port, node.Name)
	}

	return node.Link

}
