package v2

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yimuysl001/gtoolboxs/utility/myhttp"
	"strings"
)

const (
	OpenAPITitle       = `后台管理`
	OpenAPIDescription = `基于 GoFrame2.0的后台管理系统。 Enjoy 💖 `
	OpenAPIContactName = "后台管理"
	OpenAPIContactUrl  = "http://www.g-fast.cn"
	SwaggerUITemplate  = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" href="{SwaggerUIDocNamePlaceHolder}/doc/favicon.ico" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title></title>
    <script type="module" crossorigin src="{SwaggerUIDocNamePlaceHolder}/doc/webjars/js/doc-7814a93f.js"></script>
    <link rel="stylesheet" href="{SwaggerUIDocNamePlaceHolder}/doc/webjars/css/doc-e469198e.css">
  </head>
  <body>
    <div id="app"></div>
  </body>
</html>
`
)

func Swagger(ctx context.Context, s *ghttp.Server) {
	swaggerPath := g.Cfg().MustGet(ctx, "server.swaggerPath").String()
	if swaggerPath == "" {
		return
	}

	s.BindHookHandler(swaggerPath+"/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		content := gstr.ReplaceByMap(SwaggerUITemplate, map[string]string{
			`{SwaggerUIDocUrl}`:             g.Cfg().MustGet(ctx, "server.openapiPath").String(),
			`{SwaggerUIDocNamePlaceHolder}`: gstr.TrimRight(fmt.Sprintf(`//%s`, r.Host)),
		})
		r.Response.Write(content)
		r.ExitAll()
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		AddApiDocRouter(group)
	})

	openapi := s.GetOpenApi()
	openapi.Info = goai.Info{
		Title:       OpenAPITitle,
		Description: OpenAPIDescription,
		Contact: &goai.Contact{
			Name: OpenAPIContactName,
			URL:  OpenAPIContactUrl,
		},
	}

}

func AddApiDocRouter(group *ghttp.RouterGroup) {
	group.GET("/api.json", func(r *ghttp.Request) {
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Header().Set("content-type", "application/json;charset=UTF-8")

		//r.Response.RedirectTo(APIJSON)
		r.Response.Write(RedJson(r.Context()))
	})

}

func RedJson(ctx context.Context) map[string]interface{} {
	//path := serverRoot + "/doc/api.json"
	//file, err := os.ReadFile(path)
	//if err != nil {
	//	return nil
	//}
	//m := gjson.New(file).Map()
	url := "http://127.0.0.1" + g.Cfg().MustGet(ctx, "server.address").String() + g.Cfg().MustGet(ctx, "server.openapiPath").String()
	get := myhttp.HttpGet(ctx, url, nil)
	m := gjson.New(get).Map()
	newmap := make(map[string]interface{})
	for s, i := range m {
		if !strings.EqualFold(s, "paths") {
			newmap[s] = i
			continue
		}

		paths := gjson.New(i).Map()

		pathsmap := make(map[string]interface{})
	n1:
		for s2, i2 := range paths {
			if strings.Contains(s2, "/api/v1/business") || strings.Contains(s2, "/api/v1/demo") ||
				strings.Contains(s2, "/api/v1/pub") || strings.Contains(s2, "/api/v1/system") ||
				strings.Contains(s2, "/api/v1/log") {
				continue n1
			}
			pathsmap[s2] = i2

		}
		newmap[s] = pathsmap

	}

	return newmap
}
