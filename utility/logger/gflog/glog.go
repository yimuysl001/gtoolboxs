package gflog

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/yimuysl001/gtoolboxs/utility/logger/mycode"
	"strings"
)

var logctx = gctx.New()

var Gflogger = &gflogger{
	errorlog: g.Log("loggererror").Skip(1).Line(true).Async(true),
	infolog:  g.Log().Skip(1).Line(true),
	tracelog: g.Log("loggertrace").Skip(1).Line(true).Async(true),
}

type gflogger struct {
	errorlog *glog.Logger //错误日志
	infolog  *glog.Logger //一般日志
	tracelog *glog.Logger //后台日志
}

func (g *gflogger) Info(v ...interface{}) {
	g.infolog.Info(logctx, v...)

}

func (g *gflogger) InfoCtx(ctx context.Context, v ...interface{}) {
	g.infolog.Info(ctx, v...)
}

func (g *gflogger) Trace(v ...interface{}) {
	g.tracelog.Info(logctx, v...)
}

func (g *gflogger) TraceCtx(ctx context.Context, v ...interface{}) {
	g.tracelog.Info(ctx, v...)
}

func (g *gflogger) Error(v ...interface{}) {
	g.errorlog.Error(logctx, v...)
}

func (g *gflogger) ErrorCtx(ctx context.Context, v ...interface{}) {
	g.errorlog.Error(ctx, v...)
}

func (g *gflogger) IfError(err interface{}, v ...interface{}) {
	if err == nil {
		return
	}
	if v == nil || len(v) < 1 {
		g.errorlog.Error(logctx, err)
		return
	}
	var vs = make([]interface{}, 0)
	vs = append(vs, v...)
	vs = append(vs, err)

	g.errorlog.Error(logctx, vs...)
}

func (g *gflogger) IfErrorCtx(ctx context.Context, err interface{}, v ...interface{}) {
	if err == nil {
		return
	}
	if v == nil || len(v) < 1 {
		g.errorlog.Error(ctx, err)
		return
	}
	var vs = make([]interface{}, 0)
	vs = append(vs, v...)
	vs = append(vs, err)

	g.errorlog.Error(ctx, vs...)
}

func (g *gflogger) PanicCtx(ctx context.Context, text string, codes ...gcode.Code) {
	code := mycode.CodeNil
	if codes != nil && len(codes) > 0 && codes[0] != nil {
		code = codes[0]
	}
	err := gerror.WrapCode(code, fmt.Errorf("%v", text))

	g.errorlog.Error(ctx, err)

	panic(err)
}

func (g *gflogger) PanicErrorCodeCtx(ctx context.Context, err interface{}, codes ...gcode.Code) {
	if err == nil {
		return
	}
	code := mycode.CodeNil
	if codes != nil && len(codes) > 0 && codes[0] != nil {
		code = codes[0]
	}
	nerr := gerror.WrapCode(code, fmt.Errorf("%v", err))
	g.errorlog.Error(ctx, nerr)

	panic(nerr)

}

func (g *gflogger) PanicErrorCtx(ctx context.Context, err interface{}, v ...interface{}) {
	if err == nil {
		return
	}
	if v == nil || len(v) < 1 {
		g.errorlog.Error(ctx, err)
		panic(gerror.New(fmt.Sprintf("%s", err)))
	}

	var vs = make([]interface{}, 0)
	vs = append(vs, v...)
	vs = append(vs, err)

	g.errorlog.Error(ctx, vs...)

	var sb strings.Builder
	for i := 0; i < len(vs); i++ {
		sb.WriteString("%s \n")
	}

	panic(gerror.New(fmt.Sprintf(sb.String(), vs...)))

}
