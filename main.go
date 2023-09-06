package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func main() {

	//fmt.Println(pongo.GetPongoTag())

	// 原始SQL查询字符串
	sqlQuery := `SELECT "id","name","size","path","full_path","mime_type","source","describe","md5","created_by","updated_by","created_at","updated_at","deleted_at" FROM "big_file" WHERE "deleted_at" IS NULL ORDER BY "created_at" DESC LIMIT 0,10`

	// 使用正则表达式替换字段名称
	re := regexp.MustCompile(`"([^"]+)"`)
	replacedQuery := re.ReplaceAllString(sqlQuery, `"[$1]"`)

	fmt.Println(replacedQuery)

}

// getFunctionComment 获取函数的文档注释
func getFunctionComment(funcType reflect.Type) string {
	// 获取函数名称
	//funcName := funcType.Name()

	// 获取函数的定义代码
	funcDef := fmt.Sprintf("%s\n", funcType)

	// 从函数定义中提取注释
	commentIndex := strings.Index(funcDef, "//")
	if commentIndex != -1 {
		return strings.TrimSpace(funcDef[commentIndex+2:])
	}

	return ""
}
