package pongo

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

// DoPlugin
//
//	@Description: 一些简单的外部插件处理
//	@param ctx
//	@param path
//	@param data
//	@return map[string]interface{}
func DoPlugin(ctx context.Context, path string, data map[string]interface{}) map[string]interface{} {
	out, err := ParseContentFile(path, data)
	logger.Logger.PanicErrorCtx(ctx, err)
	indexData := IndexData(ctx, out)
	return gjson.New(indexData).Map()

}
