package soapdemo

import (
	"context"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

type Model struct {
	Names   string                                                       //方法名
	DoThing func(ctx context.Context, CYWBM, str string) (string, error) //数据处理具体方法
}
type ModelRequest struct {
	Xml string
}

type ModelResponse struct {
	ResultInfo string
}

// Name todo 方法名
func (u Model) Name() string {
	return u.Names
}

// ReqStruct todo  接收数据
func (u Model) ReqStruct() interface{} {
	return ModelRequest{}
}

// RespStruct todo 返回数据
func (u Model) RespStruct() interface{} {
	return ModelResponse{}
}

func (u Model) Do(ctx context.Context, req interface{}, resp interface{}) error {
	var err error
	re := req.(*ModelRequest)
	res := resp.(*ModelResponse)
	logger.Logger.TraceCtx(ctx, "=================", u.Name(), "开始=========================")
	res.ResultInfo, err = u.DoThing(ctx, u.Name(), re.Xml)
	if err != nil {
		logger.Logger.ErrorCtx(ctx, err.Error())
		res.ResultInfo = err.Error()
	}
	//logger.TraceCtx(ctx, res.ResultInfo)
	logger.Logger.TraceCtx(ctx, u.Name()+"出参：", res.ResultInfo)
	logger.Logger.TraceCtx(ctx, "=================", u.Name(), "结束=========================")
	return nil
}
