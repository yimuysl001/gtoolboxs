package soapdemo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yimuysl001/gtoolboxs/utility/webservice/soap11/dataxml"
)

type WebInter interface {
	//SetMethodName(namespace string)
	GetMethodName() string //注册方法名
	//SetNameSpace(namespace string)
	GetNameSpace() string //注册路由url地址 与url空间
	SetModel(m ...Model)  //注册业务处理逻辑
	GetModels() []Model   //业务处理逻辑
}

type Webmodel struct {
	Models     []Model //数据处理实例
	MethodName string
	NameSpace  string // http://ip:端口
}

func (w *Webmodel) GetMethodName() string {
	return w.MethodName
}

func (w *Webmodel) GetNameSpace() string {
	return w.NameSpace
}

func (w *Webmodel) SetModel(models ...Model) {
	for _, model := range models {
		if model.Name() == "" {
			panic("注册业务没有业务名称")
		}
		w.Models = append(w.Models, model)
	}
}
func (w *Webmodel) ModelAppend(models ...Model) {

	for _, model := range models {
		if model.Name() == "" {
			panic("注册业务没有业务名称")
		}
		w.Models = append(w.Models, model)
	}

}
func (w *Webmodel) GetModels() []Model {
	return w.Models
}

func WebService(in WebInter, prefix string) func(group *ghttp.RouterGroup) {
	return func(group *ghttp.RouterGroup) {
		//group.Middleware(mid.AuthMiddleware)
		methodname := in.GetMethodName()
		my := dataxml.NewServer(methodname, in.GetNameSpace()+prefix+"/"+methodname)
		models := in.GetModels()
		if len(models) == 0 {
			panic("未设置方法实例")
		}
		for _, model := range models {
			if model.DoThing == nil {
				panic(model.Names + "未设置处理方法")
			}
			my.RegisterMethod(model)
		}
		group.ALL("/"+methodname, my.Handler)
	}
}

func NewWeb(NameSpace, cywzl string) *Webmodel {
	return &Webmodel{
		MethodName: cywzl,
		NameSpace:  NameSpace,
	}
}

func GetWebModelTest(namespase string, MethodName string) Webmodel {

	var webmodel = Webmodel{
		NameSpace:  namespase, //"http://" + consts.YxConfig.NameSpace,
		MethodName: MethodName,
	}
	webmodel.Models = append(webmodel.Models, Model{
		Names: "Test",
		DoThing: func(ctx context.Context, CYWBM, str string) (string, error) {
			return "测试成功：" + str, nil
		},
	})
	return webmodel

}

func WebServiceDemo(prefix, cywlx string) func(group *ghttp.RouterGroup) {

	namespace := "http://127.0.0.1" + g.Cfg().MustGet(gctx.New(), "server.address").String()
	webmodel := NewWeb(namespace, cywlx)

	webmodel.ModelAppend(Model{
		Names:   "Test",
		DoThing: Test,
	})
	return WebService(webmodel, prefix)
}
func Test(ctx context.Context, CYWBM, str string) (xmls string, err error) {
	return CYWBM + str + "调用成功", nil
}
