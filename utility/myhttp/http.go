package myhttp

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"io"
)

// http post 请求
func HttpPost(ctx context.Context, url string, head map[string]string, datas ...interface{}) string {

	return HttpWebs(ctx, "POST", url, head, datas...)
}

func HttpGet(ctx context.Context, url string, head map[string]string, datas ...interface{}) string {
	return HttpWebs(ctx, "GET", url, head, datas...)
}

func HttpsWeb(ctx context.Context, method string, url string, head map[string]string, datas ...interface{}) {
	logger.Logger.Trace(ctx, "请求头：", head)
	logger.Logger.Trace(ctx, "请求地址：", url)
	logger.Logger.Trace(ctx, "数据提交：", datas)
}

func HttpWebs(ctx context.Context, method string, url string, head map[string]string, datas ...interface{}) string {
	logger.Logger.Trace(ctx, "请求头：", head)
	logger.Logger.Trace(ctx, "请求地址：", url)
	logger.Logger.Trace(ctx, "数据提交：", datas)

	var CL = gclient.New()
	for k, v := range head {
		CL.SetHeader(k, v)
	}

	post, err := CL.DoRequest(ctx, method, url, datas...)
	logger.Logger.PanicErrorCtx(ctx, err)
	if post == nil {
		return ""
	}

	defer post.Body.Close()
	data, err := io.ReadAll(post.Body)
	logger.Logger.PanicErrorCtx(ctx, err)
	//if err != nil {
	//	mylog.Error(ctx, "请求地址：", url)
	//	mylog.Error(ctx, "请求地址：", err)
	//	return ""
	//}
	return string(data)
}
