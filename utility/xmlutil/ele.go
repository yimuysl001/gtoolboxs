package xmlutil

import (
	"github.com/beevik/etree"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"strings"
)

// root处理方法
func TraverseXmlRoot(root *etree.Element, froot func(root *etree.Element) bool) {
	flag := froot(root)
	if !flag {
		for _, element := range root.ChildElements() {
			TraverseXmlRoot(element, froot)
		}
	}

}

// 根据xpath给节点赋值兼容@属性的模式
func SetElementValue(ROOT *etree.Element, xmlpath string, value string) {

	if strings.Contains(xmlpath, "/@") {
		sp := strings.Split(xmlpath, "/@")
		xmlpathElement := sp[0]
		xmlpathAttr := sp[1]
		if xnode := ROOT.FindElement(xmlpathElement); xnode != nil {
			if attr := xnode.SelectAttr(xmlpathAttr); attr != nil {
				attr.Value = value
			} else {
				xnode.CreateAttr(xmlpathAttr, value)
			}
		}
	} else {
		if xnode := ROOT.FindElement(xmlpath); xnode != nil {
			xnode.SetText(value)
		}
	}
}

func DelAttr(ROOT *etree.Element, key string) {
	if attr := ROOT.SelectAttr(key); attr != nil {
		ROOT.RemoveAttr(key)
	}
}

// 根据xpath给节点赋值兼容@属性的模式
func GetElementValue(ROOT *etree.Element, xmlpath string) string {
	MSG := ""
	if strings.Contains(xmlpath, "/@") {
		sp := strings.Split(xmlpath, "/@")
		xmlpathElement := sp[0]
		xmlpathAttr := sp[1]
		if xnode := ROOT.FindElement(xmlpathElement); xnode != nil {
			if attr := xnode.SelectAttr(xmlpathAttr); attr != nil {
				MSG = attr.Value
			}
		}
	} else {
		if xnode := ROOT.FindElement(xmlpath); xnode != nil {
			MSG = xnode.Text()
		}
	}
	return strings.TrimSpace(MSG)
}

func GetElement(ROOT *etree.Element, xmlpath string) *etree.Element {
	return ROOT.FindElement(xmlpath)
}

func GetElements(ROOT *etree.Element, xmlpath string) []*etree.Element {
	return ROOT.FindElements(xmlpath)
}

// xml 解析
func GetROOT(bodys string) *etree.Element {
	body := ""
	split := strings.Split(bodys, "<![CDATA[")
	if len(split) > 1 {
		body = strings.Split(split[1], "]]>")[0]
	} else {
		body = bodys
	}
	//seelog.Info(body)
	docWrite := etree.NewDocument()
	if err := docWrite.ReadFromString(body); err == nil {
		root := docWrite.Root()
		return root
	}
	return nil
}
func GetElementRoot(bodys string) *etree.Element {

	docWrite := etree.NewDocument()
	if err := docWrite.ReadFromString(bodys); err == nil {
		root := docWrite.Root()
		return root
	} else {
		panic(err)
		//mylog.Error(context.TODO(), "xml解析失败", err)
	}
	return nil
}

func RecordToXml(vs gdb.Record, rootTag ...string) string {
	var strmap = make(map[string]string)
	for k, v := range vs {
		strmap[k] = v.String()
	}
	//Mylog.MyInfo().Println("strmap：", strmap)
	xmlString := gjson.New(strmap).MustToXmlIndentString(rootTag...)
	//Mylog.MyInfo().Println("xmlString：", xmlString)
	return xmlString
}

func ResultToXml(vs gdb.Result, rootTag ...string) string {
	var sb strings.Builder

	for _, v := range vs {
		if len(rootTag) > 0 {
			sb.WriteString(RecordToXml(v, rootTag[0]))
		} else {
			sb.WriteString(RecordToXml(v))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
