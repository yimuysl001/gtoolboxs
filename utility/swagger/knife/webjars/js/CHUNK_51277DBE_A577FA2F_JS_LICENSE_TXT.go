package js

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/utils"
)

const (
	CHUNK_51277DBE_A577FA2F_JS_LICENSE_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-51277dbe.a577fa2f.js.LICENSE.txt"
	// 文件内容的16进制表示
	CHUNK_51277DBE_A577FA2F_JS_LICENSE_TXT_HEX_CONTENT = `2f2a210a202a20636c6970626f6172642e6a732076322e302e360a202a2068747470733a2f2f636c6970626f6172646a732e636f6d2f0a202a200a202a204c6963656e736564204d495420c2a9205a656e6f20526f6368610a202a2f0a`
)

func AddRouterOfChunk51277dbeA577fa2fJsLICENSETxt(router *ghttp.RouterGroup) {

	utils.GetOther(router, CHUNK_51277DBE_A577FA2F_JS_LICENSE_TXT_RELATIVE_PATH, CHUNK_51277DBE_A577FA2F_JS_LICENSE_TXT_HEX_CONTENT)

}
