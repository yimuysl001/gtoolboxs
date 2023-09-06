package dataxml

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"io"

	"github.com/gogf/gf/v2/net/ghttp"
	"reflect"
	"strings"
)

// IServer soap 服务
type IServer interface {

	// RegisterMethod 注册服务方法
	RegisterMethod(act Method) error

	// Handler http handler
	Handler(r *ghttp.Request)
}

type Method interface {
	// Name 名称
	Name() string
	// ReqStruct 请求结构体
	ReqStruct() interface{}
	// RespStruct 结果结构体
	RespStruct() interface{}
	// Do 处理请求
	Do(ctx context.Context, req interface{}, resp interface{}) error
}

type server struct {
	name      string            `description:"服务名称"`
	namespace string            `description:"wsdl文件的自主命名空间"`
	wsdl      *WsdlDefinitions  `description:"wsdl对象"`
	wsdlBytes []byte            `description:"wsdl序列化信息"`
	methods   map[string]Method `description:"服务方法"`
}

// NewServer 创建一个soap服务体
func NewServer(name string, namespace string) IServer {
	logger.Logger.Info("webservice:", namespace+"?wsdl")
	s := &server{
		name:      name,
		namespace: namespace,
		methods:   map[string]Method{},
	}
	s.buildWsdl()
	return s
}

func (s *server) buildWsdl() {
	def := &WsdlDefinitions{
		Tns:      s.namespace,
		TargetNs: s.namespace,
		Soap:     "http://schemas.xmlsoap.org/wsdl/soap/",
		SoapEnv:  "http://schemas.xmlsoap.org/soap/envelope/",
		Wsdl:     "http://schemas.xmlsoap.org/wsdl/",
		Xsd:      "http://www.w3.org/2001/XMLSchema",
		Xsi:      "http://www.w3.org/2001/XMLSchema-instance",
	}
	sch := XsdSchema{
		TargetNamespace: s.namespace,
		Import: []XsdImport{
			{Namespace: "http://schemas.xmlsoap.org/soap/encoding/"},
			{Namespace: "http://schemas.xmlsoap.org/wsdl/"}},
	}
	def.Types.Schemas = append(def.Types.Schemas, sch)

	def.PortType.Name = s.name + "PortType"

	def.Binding.Name = s.name + "Binding"
	def.Binding.Type = "tns:" + def.PortType.Name
	def.Binding.SoapBinding.Style = "rpc"
	def.Binding.SoapBinding.Transport = "http://schemas.xmlsoap.org/soap/http"

	def.Service.Name = s.name
	def.Service.Port = WsdlServicePort{
		Name:    s.name + "Port",
		Binding: "tns:" + def.Binding.Name,
		Address: WsdlServiceAddress{Location: s.namespace},
	}
	s.wsdl = def
}

func (s *server) RegisterMethod(m Method) error {
	s.wsdlBytes = nil
	// 方法名
	tname := m.Name()
	if _, ok := s.methods[tname]; ok {
		logger.Logger.Error("方法重复注册:" + tname)
		return errors.New("方法重复注册:" + tname)
	}

	// message
	erro := s.parseMessage(tname+"Request", m.ReqStruct())
	if erro != nil {
		logger.Logger.Error(erro)
		return erro
	}
	erro = s.parseMessage(tname+"Respone", m.RespStruct())
	if erro != nil {
		logger.Logger.Error(erro)
		return erro
	}
	logger.Logger.Info("已注册服务：", tname)
	s.regWsdl(tname)
	s.methods[tname] = m
	return nil
}

// 解析参数并转化为对应的wsdl message
func (s *server) parseMessage(name string, st interface{}) error {
	msg := WsdlMessage{Name: name}
	retype := reflect.TypeOf(st)
	if retype.Kind() == reflect.Ptr {
		retype = retype.Elem()
	} else if retype.Kind() == reflect.Slice {
		retype = retype.Elem()
	}
	if retype.Kind() != reflect.Struct {
		return errors.New(name + "必须是一个结构体")
	}
	// 遍历结构体参数列表
	for i := 0; i < retype.NumField(); i++ {
		name, _ := getTagsInfo(retype.Field(i))
		ik := retype.Field(i).Type.Kind()
		ts, erro := checkBaseTypeKind(ik)
		// 如果非基本类型则转为自有命名空间的自定义类型
		if erro != nil {
			ts = "tns:" + name + ik.String()
			tp := reflect.New(retype.Field(i).Type).Elem().Interface()
			_ = s.parseMessage(name+ik.String(), tp)
		}
		msg.Part = append(msg.Part, WsdlPart{Name: name, Type: ts})
	}
	s.wsdl.Message = append(s.wsdl.Message, msg)
	return nil
}

func (s *server) regWsdl(methodName string) {
	// portype
	op := WsdlPortTypeOperation{
		Name:   methodName,
		Input:  WsdlPortTypeOperationMessage{Message: "tns:" + methodName + "Request"},
		Output: WsdlPortTypeOperationMessage{Message: "tns:" + methodName + "Respone"},
	}
	s.wsdl.PortType.Operations = append(s.wsdl.PortType.Operations, op)
	// binding
	soapio := WsdlSoapBodyIO{SoapBody: WsdlSoapBody{Use: "encoded", EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/"}}
	bindop := WsdlBindingOperation{
		Name: methodName, Input: soapio,
		Output: soapio,
		SoapOperation: WsdlSoapOperation{
			Style: "rpc", SoapAction: s.wsdl.Tns + "/" + methodName,
		},
	}
	s.wsdl.Binding.Operations = append(s.wsdl.Binding.Operations, bindop)
}

func (s *server) Handler(r *ghttp.Request) {
	//t:=time.Now()
	w := r.Response
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Header().Set("Accpet", "text/xml")
	w.Write([]byte(xml.Header))
	if r.Method == "GET" {
		//InsertTable(r,t,"")
		// 网址带参数wsdl则显示wsdl文件
		if strings.EqualFold("wsdl", r.URL.RawQuery) {
			if s.wsdlBytes == nil {
				b, erro := xml.Marshal(s.wsdl)
				if erro != nil {
					w.Write([]byte("err:" + erro.Error()))
					return
				}
				s.wsdlBytes = b
			}
			w.Write(s.wsdlBytes)
		} else {
			// 其他情况返回一个提示信息
			w.Write([]byte(r.URL.String()))
		}
		return
	}

	// post请求则处理接受的xml
	// 读取post的body信息
	b, erro := io.ReadAll(r.Body)
	if erro != nil || len(b) < 1 {
		s.handlerResponseError("读取body出错", r, 500)
		return
	}

	defer r.Body.Close()
	//r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	// 转化为Envelope对象
	env := Envelope{}
	errxml := xml.Unmarshal(b, &env)
	if errxml != nil {
		s.handlerResponseError("xml 解析失败", r)
		return
	}
	// 解析请求的方法名字
	var startEle *xml.StartElement
	reader := bytes.NewReader(env.Body.Content)
	if reader == nil {
		s.handlerResponseError("reader 解析失败", r)
		return
	}
	de := xml.NewDecoder(reader)

	for {
		t, erro := de.Token()
		if erro != nil {
			break
		}
		if x, ok := t.(xml.StartElement); ok {
			startEle = &x
			break
		}
	}
	if startEle == nil {
		s.handlerResponseError("接受到的data无效", r)
		return
	}
	s.request(r, de, startEle)
}

func (s *server) handlerResponseError(errMsg string, r *ghttp.Request, errCode ...int64) {
	fault := ErrResult{
		Code:      500,
		ErrString: errMsg,
	}
	if len(errCode) > 0 {
		fault.Code = errCode[0]
	}
	data, _ := xml.Marshal(fault)
	b, _ := xml.Marshal(NewEnvelope(data))
	r.Response.Write(b)
}
func (s *server) request(r *ghttp.Request, de *xml.Decoder, startEle *xml.StartElement) {

	mname := startEle.Name.Local

	m, has := s.methods[mname]
	if !has {
		s.handlerResponseError("["+mname+"]不存在", r)
		return
	}
	// 解析入参
	params := m.ReqStruct()
	reqType := reflect.TypeOf(params)
	params = reflect.New(reqType).Elem().Addr().Interface()
	erro := de.DecodeElement(params, startEle)
	if erro != nil {
		s.handlerResponseError("参数错误:"+erro.Error(), r)
		return
	}
	resp := m.RespStruct()
	retype := reflect.TypeOf(resp)
	resp = reflect.New(retype).Elem().Addr().Interface()
	err := m.Do(r.Context(), params, resp)
	if err != nil {
		s.handlerResponseError("请求失败:"+err.Error(), r)
		return
	}
	bs, e := xml.Marshal(resp)
	if e != nil {
		s.handlerResponseError("结果序列化错误:"+e.Error(), r)
		return
	}
	b, _ := xml.Marshal(NewEnvelope(bs))
	r.Response.Write(b)
}
