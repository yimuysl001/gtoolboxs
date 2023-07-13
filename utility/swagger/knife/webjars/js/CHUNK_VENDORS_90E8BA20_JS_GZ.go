package js

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/utils"
)

const (
	CHUNK_VENDORS_90E8BA20_JS_GZ_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-vendors.90e8ba20.js.gz"
	// 文件内容的16进制表示
)

func AddRouterOfChunkVendors90e8ba20JsGz(router *ghttp.RouterGroup) {

	utils.GetOther(router, CHUNK_VENDORS_90E8BA20_JS_GZ_RELATIVE_PATH, CHUNK_VENDORS_90E8BA20_JS_GZ_HEX_CONTENT)

}