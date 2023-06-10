package rabbitmq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/streadway/amqp"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"sync"
	"time"
)

type RabbitMQ struct {
	sync.RWMutex
	idname        string
	url           string
	iThread       int    //消费线程处理数
	autoReconnect bool   // 是否支持断网重连
	queueName     string //队列名称
	exchangeName  string //交换机名称
	routingKey    string //key名称
	exchangeType  string //key名称
	autoTime      int
	istatus       int
	conn          *amqp.Connection
	channel       *amqp.Channel
}

// NewRabbitMQ 创建 RabbitMQ 实例
func NewRabbitMQ(ms MqStruct, name string) (*RabbitMQ, error) {
	url := fmt.Sprintf(mqurl, ms.UsernameName, ms.PasswordName, ms.NameServer, ms.Port)

	return &RabbitMQ{
		url:           url,
		idname:        name,
		iThread:       ms.IThread,
		autoReconnect: ms.AutoReconnect,
		queueName:     ms.QueueNmae,
		exchangeName:  ms.ExchangeName,
		routingKey:    ms.RoutingKey,
		exchangeType:  ms.ExchangeType,
		autoTime:      ms.AutoTime,
		conn:          nil,
		channel:       nil,
	}, nil
}

// 连接
func (r *RabbitMQ) connect() error {
	r.RLock()
	defer r.RUnlock()
	if r.istatus == 3 {
		//已成功连接
		return nil
	}
	var err error
	r.conn, err = amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}

	// 用于检查队列是否存在,已经存在不需要重复声明
	// 队列不存在,声明队列
	// name:队列名称;durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;noWait:是否非阻塞,
	// true为是,不等待RMQ返回信息;args:参数,传nil即可;exclusive:是否设置排他
	_, err = r.channel.QueueDeclare(r.queueName, true, false, false, true, nil)
	if err != nil {
		logger.Logger.Error("QueueDeclare 失败：", err)
	}

	// 注册交换机
	// name:交换机名称,kind:交换机类型,durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;
	// noWait:是否非阻塞, true为是,不等待RMQ返回信息;args:参数,传nil即可; internal:是否为内部
	err = r.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		logger.Logger.Error("QueueDeclare 失败：", err)
	}

	r.istatus = 3
	return nil

}

// Publish 发布消息
func (r *RabbitMQ) Publish(body []byte) (err error) {
	defer func() {
		if er := recover(); er != nil {
			r.istatus = 0
			logger.Logger.IfError(er)
			err = gerror.Newf("%s", er)
		} else if err != nil {
			r.istatus = 0
			logger.Logger.IfError(er)
			err = gerror.Newf("%s", er)
		}

	}()

	if r.istatus != 3 {
		logger.Logger.IfError(r.connect())
	}
	// 绑定任务
	err = r.channel.QueueBind(r.queueName, r.routingKey, r.exchangeName, true, nil)
	if err != nil {
		return err
	}
	return r.channel.Publish(
		r.exchangeName,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)

}

//func (r *RabbitMQ) consume(ctx context.Context, handler func(ctx2 context.Context, body []byte) error) error {
//	logger.Logger.InfoCtx(ctx, "线程数：", r.iThread)
//	//fmt.Println("iThread:", r.iThread)
//	// 获取消费通道,确保rabbitMQ一个一个发送消息
//	err := r.channel.Qos(r.iThread, 0, true)
//	msgs, err := r.channel.Consume(
//		r.queueName,
//		"",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		return err
//	}
//	go func() {
//		for d := range msgs {
//			err := handler(ctx, d.Body)
//			if err != nil {
//				logger.Logger.ErrorCtx(ctx, "Failed to handle message:", err)
//				d.Ack(true)
//			} else {
//				d.Ack(false)
//
//			}
//
//			//}()
//
//			//rpool.Add(ctx, func(ctx context.Context) {
//			//	err := handler(ctx, finild)
//			//	if err != nil {
//			//		logger.Logger.ErrorCtx(ctx, "Failed to handle message:", err)
//			//	} else {
//			//		d.Ack(false)
//			//	}
//			//})
//
//		}
//	}()
//
//	return nil
//}

// Consume 消费消息
func (r *RabbitMQ) Consume(ctx context.Context, handler func(ctx2 context.Context, body []byte) error) (err error) {
	defer func() {
		if er := recover(); er != nil || err != nil {
			logger.Logger.ErrorCtx(ctx, err, r.idname+"监听出错：")
			logger.Logger.IfErrorCtx(ctx, er, r.idname+"监听未知出错：")
			r.istatus = 0
			r.restart(ctx, handler)
		}
	}()

	if r.iThread <= 0 {
		r.iThread = 1
	}
	logger.Logger.InfoCtx(ctx, r.idname+"开启线程数：", r.iThread)
	//fmt.Println("iThread:", r.iThread)
	// 获取消费通道,确保rabbitMQ一个一个发送消息
	err = r.channel.Qos(r.iThread, 0, true)
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	defer func() {
		g.TryCatch(ctx, func(ctx context.Context) {
			logger.Logger.PanicErrorCtx(ctx, r.conn.Close())
		}, func(ctx context.Context, exception error) {
			logger.Logger.IfErrorCtx(ctx, exception)
		})
	}()

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer g.TryCatch(ctx, func(ctx context.Context) {
		logger.Logger.PanicErrorCtx(ctx, ch.Close())
	}, func(ctx context.Context, exception error) {
		logger.Logger.IfErrorCtx(ctx, exception)
	})

	closeChan := make(chan *amqp.Error, 1)
	notifyClose := ch.NotifyClose(closeChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
	closeFlag := false
	rpool := grpool.New(r.iThread)
	//c := make(chan bool, r.iThread)
	for {
		select {
		case e := <-notifyClose:
			logger.Logger.ErrorCtx(ctx, e.Error())
			close(closeChan)
			r.restart(ctx, handler)
			closeFlag = true
		case msg := <-msgs:
			rpool.Add(ctx, func(ctx context.Context) {
				err := handler(ctx, msg.Body)
				if err != nil {
					logger.Logger.ErrorCtx(ctx, "Failed to handle message:", err)
					msg.Ack(true)
				} else {
					msg.Ack(false)
				}
			})
		}
		if closeFlag {
			break
		}
	}

	return nil
}

// 重启服务
func (r *RabbitMQ) restart(ctx context.Context, handler func(ctx2 context.Context, body []byte) error) {
	r.istatus = 0
	if r.autoReconnect && r.autoTime > 0 {
		logger.Logger.ErrorCtx(ctx, r.autoTime, "秒后重连")
		time.Sleep(time.Duration(r.autoTime) * time.Second)
		r.connect()
		r.Consume(ctx, handler)
	}
}

// Close 关闭连接
func (r *RabbitMQ) Close() error {
	err := r.channel.Close()
	if err != nil {
		return err
	}

	err = r.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

//// Reconnect 重新连接
//func (r *RabbitMQ) Reconnect() error {
//
//	if !r.autoReconnect {
//		return nil
//	}
//	if r.autoTime <= 0 {
//		r.autoTime = 10 //默认10秒
//	}
//
//	go func() {
//		for {
//			var err error
//			r.conn, err = amqp.Dial(r.url)
//			if err == nil {
//				time.Sleep(time.Duration(r.autoTime) * time.Second)
//				continue
//			}
//			logger.Logger.Error("Failed to connect to RabbitMQ: ", err)
//			//log.Printf("Failed to connect to RabbitMQ: %s", err)
//			r.channel, err = r.conn.Channel()
//
//			logger.Logger.Error("Failed to connect to RabbitMQ: ", err)
//		}
//	}()
//
//	return nil
//}
