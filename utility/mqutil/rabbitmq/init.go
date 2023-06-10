package rabbitmq

import (
	"context"
	"github.com/yimuysl001/gtoolboxs/utility/constsutil"
	"github.com/yimuysl001/gtoolboxs/utility/maputil"
)

// amqp://账户:密码@IP:端口/
const mqurl = "amqp://%s:%s@%s:%d/"

type MqStruct struct {
	MqType        string //默认rabbitmq
	NameServer    string //ip
	Port          int    //端口
	QueueNmae     string //队列名称
	ExchangeName  string //交换机名称
	RoutingKey    string //key名称
	ExchangeType  string
	UsernameName  string //用户名
	PasswordName  string //密码
	AutoReconnect bool   // 是否支持断网重连
	AutoTime      int    //检测断连时间间隔
	IThread       int    //消费线程处理数
}

// 本地配置
var bdmap = make(map[string]MqStruct)

func init() {
	constsutil.InitConf(context.Background(), "MQ", &bdmap)
}

func SetMq(name string, obj MqStruct) {
	maputil.Set(bdmap, name, obj)

}
