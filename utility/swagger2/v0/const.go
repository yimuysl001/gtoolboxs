package v0

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	OpenAPITitle       = `后台管理`
	OpenAPIDescription = `基于 GoFrame2.0的后台管理系统。 Enjoy 💖 `
	OpenAPIContactName = "后台管理"
	OpenAPIContactUrl  = "http://www.g-fast.cn"
	SwaggerUITemplate  = `<!DOCTYPE html>
<html lang=en>
<head>
    <meta charset=utf-8>
    <meta http-equiv=X-UA-Compatible content="IE=edge">
    <meta name=viewport content="width=device-width,initial-scale=1">
    <link rel=icon href=favicon.ico>
    <title></title>
    <link href=knife/webjars/css/chunk-51277dbe.57225f85.css rel=prefetch>
    <link href=knife/webjars/js/chunk-069eb437.371ae4fd.js rel=prefetch>
    <link href=knife/webjars/js/chunk-0fd67716.d57e2c41.js rel=prefetch>
    <link href=knife/webjars/js/chunk-2d0af44e.c299c1d4.js rel=prefetch>
    <link href=knife/webjars/js/chunk-2d0bd799.cc91c520.js rel=prefetch>
    <link href=knife/webjars/js/chunk-2d0d0b98.cb1dea78.js rel=prefetch>
    <link href=knife/webjars/js/chunk-2d0da532.dd3c929c.js rel=prefetch>
    <link href=knife/webjars/js/chunk-2d22269d.bd9173e1.js rel=prefetch>
    <link href=knife/webjars/js/chunk-3b888a65.8737ce4f.js rel=prefetch>
    <link href=knife/webjars/js/chunk-3ec4aaa8.a79d19f8.js rel=prefetch>
    <link href=knife/webjars/js/chunk-51277dbe.a577fa2f.js rel=prefetch>
    <link href=knife/webjars/js/chunk-589faee0.b24e5f3d.js rel=prefetch>
    <link href=knife/webjars/js/chunk-735c675c.76ef1019.js rel=prefetch>
    <link href=knife/webjars/js/chunk-adb9e944.b888f4bd.js rel=prefetch>
    <link href=knife/webjars/css/app.b848c085.css rel=preload as=style>
    <link href=knife/webjars/css/chunk-vendors.3f2387de.css rel=preload as=style>
    <link href=knife/webjars/js/app.2650dddf.js rel=preload as=script>
    <link href=knife/webjars/js/chunk-vendors.90e8ba20.js rel=preload as=script>
    <link href=knife/webjars/css/chunk-vendors.3f2387de.css rel=stylesheet>
    <link href=knife/webjars/css/app.b848c085.css rel=stylesheet>
</head>
<body>
<noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it to
    continue.</strong></noscript>
<div id=app></div>
<script src=knife/webjars/js/chunk-vendors.90e8ba20.js></script>
<script src=knife/webjars/js/app.2650dddf.js></script>
</body>
</html>`
)

func Swagger(ctx context.Context, s *ghttp.Server) {
	swaggerPath := g.Cfg().MustGet(ctx, "server.swaggerPath").String()
	if swaggerPath == "" {
		return
	}

	s.BindHookHandler(swaggerPath+"/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		//content := gstr.ReplaceByMap(SwaggerUITemplate, map[string]string{
		//	`{SwaggerUIDocUrl}`:             g.Cfg().MustGet(ctx, "server.openapiPath").String(),
		//	`{SwaggerUIDocNamePlaceHolder}`: gstr.TrimRight(fmt.Sprintf(`//%s`, r.Host)),
		//})
		r.Response.Write(SwaggerUITemplate)
		r.ExitAll()
	})
	s.Group("/knife", func(group *ghttp.RouterGroup) {
		AddApiDocRouter(group)
		AddSwaggerResourcesRouter(group)
	})

	//openapi := s.GetOpenApi()
	//openapi.Info = goai.Info{
	//	Title:       OpenAPITitle,
	//	Description: OpenAPIDescription,
	//	Contact: &goai.Contact{
	//		Name: OpenAPIContactName,
	//		URL:  OpenAPIContactUrl,
	//	},
	//}
}
