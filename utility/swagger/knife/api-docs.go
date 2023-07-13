package knife

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
)

const (
	APIJSON                = "/api.json"
	API_DOCS_RELATIVE_PATH = constant.ROOT_PATH + "/v2/api-docs"
)

func AddApiDocRouter(group *ghttp.RouterGroup) {
	group.GET(API_DOCS_RELATIVE_PATH, func(r *ghttp.Request) {
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Header().Set("content-type", "application/json;charset=UTF-8")
		r.Response.RedirectTo(APIJSON)
	})
}
