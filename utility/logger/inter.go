package logger

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
)

type MyLogger interface {
	// Info 需要打印控制台的日志
	Info(v ...interface{})
	// InfoCtx 需要打印控制台的日志,添加上下文
	InfoCtx(ctx context.Context, v ...interface{})
	// Trace 不打印控制台的日志
	Trace(v ...interface{})
	// TraceCtx 不打印控制台的日志,添加上下文
	TraceCtx(ctx context.Context, v ...interface{})
	// Error 错误日志
	Error(v ...interface{})
	// ErrorCtx 错误日志，添加上下文
	ErrorCtx(ctx context.Context, v ...interface{})
	// IfError 如果有错误err不为空，打印日志
	IfError(err interface{}, v ...interface{})
	// IfErrorCtx 如果有错误err不为空，打印日志，添加上下文
	IfErrorCtx(ctx context.Context, err interface{}, v ...interface{})
	// PanicCtx 自定义错误
	PanicCtx(ctx context.Context, text string, code ...gcode.Code)
	// PanicErrorCtx 如果err 不为空，打印错误日志并抛异常,添加上下文 v 补充说明
	PanicErrorCtx(ctx context.Context, err interface{}, v ...interface{})

	PanicErrorCodeCtx(ctx context.Context, err interface{}, code ...gcode.Code)
}
