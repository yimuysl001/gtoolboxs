package main

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/xmlutil"
)

func main() {
	//ole.CoInitialize(0)
	//unknown, _ := oleutil.CreateObject("Excel.Application")
	//excel, _ := unknown.QueryInterface(ole.IID_IDispatch)
	//oleutil.PutProperty(excel, "Visible", true)
	//workbooks := oleutil.MustGetProperty(excel, "Workbooks").ToIDispatch()
	//workbook := oleutil.MustCallMethod(workbooks, "Add", nil).ToIDispatch()
	//worksheet := oleutil.MustGetProperty(workbook, "Worksheets", 1).ToIDispatch()
	//cell := oleutil.MustGetProperty(worksheet, "Cells", 1, 1).ToIDispatch()
	//oleutil.PutProperty(cell, "Value", 12345)
	//
	//time.Sleep(time.Hour)
	//
	//oleutil.PutProperty(workbook, "Saved", true)
	//oleutil.CallMethod(workbook, "Close", false)
	//oleutil.CallMethod(excel, "Quit")
	//excel.Release()
	//
	//ole.CoUninitialize()

	//ole.CoInitialize(0)
	//unknown, err := oleutil.CreateObject("MyCOMComponent.MyCOMClass")
	//fmt.Println(err)
	//excel, err := unknown.QueryInterface(ole.NewGUID("{12345678-1234-1234-1234-1234567890AB}"))
	//fmt.Println(err)
	//_, err = oleutil.PutProperty(excel, "MyMethod")
	//fmt.Println(err)
	//time.Sleep(time.Hour)
	//ole.CoUninitialize()

	//limit, e := dbutil.DB().Model("test").Page(1, 10).All()
	//fmt.Println(limit, e)
	jst := `{
	"body": {
		"MS185": [
			{
				"action": "insert",
				"code": "LIS9040",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "咽拭子",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "YSZ",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS9034",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "肺泡灌洗液",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "FPGXY",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1800",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "心包积液",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "XBJY",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1755",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "脑脊液",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "NJY",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1756",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "胸水",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "XS",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS9082",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "疱液",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "PY",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1757",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "腹水",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "FS",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1763",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "痰液",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "TY",
				"stopflag": ""
			},
			{
				"action": "insert",
				"code": "LIS1748",
				"defaultSpecimen": "0",
				"idHis": "18365",
				"idLis": "LIS011093",
				"itemVersion": 1,
				"name": "全血",
				"nameHis": "19种病原体核酸多重检测",
				"nameLis": "19种病原体核酸多重检测",
				"pyCode": "QX",
				"stopflag": ""
			}
		],
		"codeSystem": null,
		"createTime": 20230811103735,
		"docImageContent": "",
		"examItems": [
			{
				"imageText": ""
			}
		],
		"msgId": "MS185",
		"msgName": "检验项目标本对照关系信息",
		"sourceSysCode": "S028",
		"targetSysCode": ""
	},
	"header": {
		"apply_unit_id": "0",
		"date_time": null,
		"exec_unit_id": "0",
		"extend_sub_id": "0",
		"hospital_id": "49557491-0",
		"msg_id": "414d51204757492e514d20202020202011619d648e6d6323",
		"order_exec_id": "0",
		"send_sys_id": "S011",
		"service_id": "MS185",
		"target_hospital": null,
		"visit_type": null
	}
}`

	m := gjson.New(jst).Map()
	xmlutil.ConvertIntToString(&m)
	logger.Logger.Info(gjson.New(m).MustToXmlString())

}
