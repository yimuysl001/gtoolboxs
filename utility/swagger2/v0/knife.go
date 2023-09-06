package v0

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strconv"
)

const (
	APIJSON                = "/api.json"
	API_DOCS_RELATIVE_PATH = "/v2/api-docs"
)

func AddApiDocRouter(group *ghttp.RouterGroup) {
	group.GET(API_DOCS_RELATIVE_PATH, func(r *ghttp.Request) {
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Header().Set("content-type", "application/json;charset=UTF-8")

		//path := serverRoot + "/knife/" + APIJSON
		//file, err := os.ReadFile(path)
		//logger.Logger.PanicErrorCtx(r.Context(), err)
		//m := gjson.New(file).Map()
		//newmap := make(map[string]interface{})
		//
		//for s, i := range m {
		//	if !strings.EqualFold(s, "paths") {
		//		newmap[s] = i
		//		continue
		//	}
		//	paths := gjson.New(i).Map()
		//
		//	pathsmap := make(map[string]interface{})
		//	//n1:
		//	for s2, i2 := range paths {
		//		//if strings.Contains("/api/v1/business", s2) || strings.Contains("/api/v1/demo", s2) ||
		//		//	strings.Contains("/api/pub/demo", s2) || strings.Contains("/api/pub/system", s2) {
		//		//	continue n1
		//		//}
		//		ncpath := make(map[string]interface{})
		//	n2:
		//		for i3, i4 := range gjson.New(i2).Map() {
		//			if strings.EqualFold("summary", i3) {
		//				ncpath[i3] = i4
		//				continue n2
		//			}
		//			pjson := gjson.New(i4)
		//			pjson.Set("operationId", guid.S())
		//
		//			ncpath[i3] = pjson.Map()
		//		}
		//		pathsmap[s2] = ncpath
		//
		//	}
		//	newmap[s] = paths
		//
		//}

		//r.Response.Write(file)

		r.Response.RedirectTo(APIJSON)
	})

}

const (
	SWAGGER_RESOURCES_CONTENT       = `[{"location":"/v2/api-docs?group=2.X版本","name":"2.X版本","swaggerVersion":"2.0","url":"/v2/api-docs?group=2.X版本"}]`
	SWAGGER_RESOURCES_RELATIVE_PATH = "/swagger-resources"
)

func AddSwaggerResourcesRouter(router *ghttp.RouterGroup) {
	router.GET(SWAGGER_RESOURCES_RELATIVE_PATH, func(r *ghttp.Request) {
		rs := []byte(SWAGGER_RESOURCES_CONTENT)
		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-type", "application/json;charset=UTF-8")
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)

	})
}
