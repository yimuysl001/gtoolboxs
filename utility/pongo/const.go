package pongo

import (
	"fmt"
	"reflect"
	"strings"
)

type PongoTag struct {
	Key      string
	Name     string
	Index    int
	FuncName string
	SetTag   string
	Tag      string
}

func GetPongoTag() []PongoTag {
	funcs := GetFuncs()
	var pts = make([]PongoTag, len(funcs))

	index := 0
	for s, a := range funcs {

		funcValue := reflect.ValueOf(a)
		// 获取函数的类型
		funcType := funcValue.Type()

		npt := PongoTag{
			Key:   s,
			Name:  s + "()",
			Index: index,
			FuncName: func() string {
				f := funcType.Name()
				if f == "" {
					f = s
				}
				return f
			}(),
			SetTag: "",
			Tag:    "",
		}
		numArgs := funcType.NumIn()
		var inargs = make([]string, numArgs)
		// 遍历函数的入参
		for i := 0; i < numArgs; i++ {
			argType := funcType.In(i)
			argTypes := strings.ReplaceAll(argType.String(), "interface {}", "any")

			inargs[i] = fmt.Sprintf(`${%d:%s}`, i+1, argTypes)

		}
		// 输出函数入参个数
		outArgs := funcType.NumOut()
		var outargs = make([]string, outArgs)
		// 遍历函数的入参
		for i := 0; i < outArgs; i++ {
			argType := funcType.Out(i)
			argTypes := strings.ReplaceAll(argType.String(), "interface {}", "any")
			outargs[i] = fmt.Sprintf(`${%d:%s}`, numArgs+i+1, argTypes)
		}

		npt.Tag = s + "(%v)"
		npt.Tag = fmt.Sprintf(npt.Tag, strings.Join(inargs, ","))

		npt.SetTag = " set  %v = %v "
		npt.SetTag = fmt.Sprintf(npt.SetTag, strings.Join(outargs, ","), npt.Tag)
		npt.SetTag = "{% " + npt.SetTag + " %}"
		pts[index] = npt
		index++
	}

	return pts

}
