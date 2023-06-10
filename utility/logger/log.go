package logger

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/yimuysl001/gtoolboxs/utility/logger/gflog"
	"go.opentelemetry.io/otel/trace"
)

const Loggerkey = "RequestId"

var Logger MyLogger

func init() {
	Logger = log()
}

func log() MyLogger {

	return gflog.Gflogger

}

func GetUuid(ctx context.Context) string {

	id := ""
	if ctx == nil {
		return guid.S()
	}

	//自定义传key
	var ctxValue interface{}
	if ctxValue = ctx.Value(Loggerkey); ctxValue == nil {
		ctxValue = ctx.Value(gctx.StrKey(gconv.String(Loggerkey)))
	}
	if ctxValue == nil {
		return guid.S()
	}
	id = gconv.String(ctxValue)
	if id != "" {
		return id
	}

	//数据请求的id
	// Tracing values.
	spanCtx := trace.SpanContextFromContext(ctx)
	if traceId := spanCtx.TraceID(); traceId.IsValid() {
		id = traceId.String()
	}
	if id != "" {
		return id
	}

	return guid.S()

}
