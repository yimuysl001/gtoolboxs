package knife

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/constant"
	"github.com/yimuysl001/gtoolboxs/utility/swagger/utils"
)

const (
	SWAGGER_RESOURCES_CONTENT       = `[{"location":"/v2/api-docs?group=2.X版本","name":"2.X版本","swaggerVersion":"2.0","url":"/v2/api-docs?group=2.X版本"}]`
	SWAGGER_RESOURCES_RELATIVE_PATH = constant.ROOT_PATH + "/swagger-resources"
)

func AddSwaggerResourcesRouter(router *ghttp.RouterGroup) {
	utils.GetJson(router, SWAGGER_RESOURCES_RELATIVE_PATH, SWAGGER_RESOURCES_CONTENT)
}
