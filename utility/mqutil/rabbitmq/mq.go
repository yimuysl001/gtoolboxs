package rabbitmq

import (
	"context"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/maputil"
)

const defaultname = "default"

// mq连接
var mqmap = make(map[string]*RabbitMQ)

func MQ(name ...string) *RabbitMQ {
	n := defaultname
	if name != nil && len(name) > 0 && name[0] != "" {
		n = name[0]
	}
	return maputil.GetOrSetFunc(mqmap, n, func() *RabbitMQ {
		mqStruct, ok := bdmap[n]
		if !ok {
			logger.Logger.PanicCtx(context.Background(), n+" 未找到相应的配置")
		}
		mq, err := NewRabbitMQ(mqStruct, n)
		logger.Logger.PanicErrorCtx(context.Background(), err)
		logger.Logger.PanicErrorCtx(context.Background(), mq.connect())
		return mq
	})

}

func NewMQ(name ...string) *RabbitMQ {
	n := defaultname
	if name != nil && len(name) > 0 && name[0] != "" {
		n = name[0]
	}
	mqStruct, ok := bdmap[n]
	if !ok {
		logger.Logger.PanicCtx(context.Background(), n+" 未找到相应的配置")
	}
	mq, err := NewRabbitMQ(mqStruct, n)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	logger.Logger.PanicErrorCtx(context.Background(), mq.connect())
	return mq
}
