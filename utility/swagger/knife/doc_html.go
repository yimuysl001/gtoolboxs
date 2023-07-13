package knife

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/utils"
)

const (
	DOC_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/doc.html"
	// 文件内容的16进制表示
	DOC_HTML_HEX_CONTENT = `<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content="IE=edge"><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=favicon.ico><title></title><link href=webjars/css/chunk-51277dbe.57225f85.css rel=prefetch><link href=webjars/js/chunk-069eb437.371ae4fd.js rel=prefetch><link href=webjars/js/chunk-0fd67716.d57e2c41.js rel=prefetch><link href=webjars/js/chunk-2d0af44e.c299c1d4.js rel=prefetch><link href=webjars/js/chunk-2d0bd799.cc91c520.js rel=prefetch><link href=webjars/js/chunk-2d0d0b98.cb1dea78.js rel=prefetch><link href=webjars/js/chunk-2d0da532.dd3c929c.js rel=prefetch><link href=webjars/js/chunk-2d22269d.bd9173e1.js rel=prefetch><link href=webjars/js/chunk-3b888a65.8737ce4f.js rel=prefetch><link href=webjars/js/chunk-3ec4aaa8.a79d19f8.js rel=prefetch><link href=webjars/js/chunk-51277dbe.a577fa2f.js rel=prefetch><link href=webjars/js/chunk-589faee0.b24e5f3d.js rel=prefetch><link href=webjars/js/chunk-735c675c.76ef1019.js rel=prefetch><link href=webjars/js/chunk-adb9e944.b888f4bd.js rel=prefetch><link href=webjars/css/app.b848c085.css rel=preload as=style><link href=webjars/css/chunk-vendors.3f2387de.css rel=preload as=style><link href=webjars/js/app.2650dddf.js rel=preload as=script><link href=webjars/js/chunk-vendors.90e8ba20.js rel=preload as=script><link href=webjars/css/chunk-vendors.3f2387de.css rel=stylesheet><link href=webjars/css/app.b848c085.css rel=stylesheet></head><body><noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=webjars/js/chunk-vendors.90e8ba20.js></script><script src=webjars/js/app.2650dddf.js></script></body></html>`
)

func AddRouterOfDocHtml(router *ghttp.RouterGroup) {
	utils.GetHtml(router, DOC_HTML_RELATIVE_PATH, DOC_HTML_HEX_CONTENT)
}
