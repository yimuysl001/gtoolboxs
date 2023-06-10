package mqutil

import (
	"context"
	"github.com/yimuysl001/gtoolboxs/utility/mqutil/rabbitmq"
)

type MqInter interface {
	// Publish 发布消息
	Publish(body []byte) error
	// Consume 消费消息
	Consume(ctx context.Context, handler func(ctx2 context.Context, body []byte) error) (err error)
}

// MQ 数据发送使用此方法
func MQ(name ...string) MqInter {
	switch "" {
	default: //暂时只支持 rabbitmq
		return rabbitmq.MQ(name...)

	}

}

// New 数据监听最好使用此方法
func New(name ...string) MqInter {
	switch "" {
	default: //暂时只支持 rabbitmq
		return rabbitmq.NewMQ(name...)
	}

}
