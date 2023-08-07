package pongo

import (
	"context"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/stringutil"
	"regexp"
	"strings"
)

func IndexDataMust(ctx context.Context, template string) string {
	template = strings.TrimSpace(template)
	logger.Logger.TraceCtx(ctx, "IndexData原始出参：", template)

	if template == "" {
		return template
	}
	t := strings.Clone(template)
	err := g.Try(ctx, func(ctx context.Context) {
		//xml
		if template[0] == '<' {
			template = stringutil.IndentXml(ctx, template)
			//json
		} else if template[0] == '{' || template[0] == '[' {
			template = stringutil.IndentJson(ctx, replaceJson(ctx, template))
		}
	})
	if err != nil {
		template = t
	}
	return template
}

func IndexData(ctx context.Context, template string) string {

	logger.Logger.TraceCtx(ctx, "IndexData原始出参：", template)
	template = strings.TrimSpace(template)
	if template == "" {
		return template
	}
	//xml
	if template[0] == '<' {
		template = stringutil.IndentXml(ctx, template)
		//json
	} else {
		template = stringutil.IndentJson(ctx, replaceJson(ctx, template))
	}

	return template
}

func replaceJson(ctx context.Context, template string) string {
	var jsoncomma = `,\s*\]`
	compile1 := regexp.MustCompile(jsoncomma)
	template = compile1.ReplaceAllString(template, "]")

	var yhstr = `"(?s:(.*?))"`
	compile := regexp.MustCompile(yhstr)
	allString := compile.FindAllString(template, -1)
	//
	splits := compile.Split(template, -1)

	var sb strings.Builder
	sb.WriteString(splits[0])
	for i, split := range allString {
		split = strings.ReplaceAll(split, "\n", "\\n")
		split = strings.ReplaceAll(split, "\r", "\\r")
		sb.WriteString(split)
		sb.WriteString(splits[i+1])
	}

	return sb.String()

}

func ParseContent(str string, data pongo2.Context) (out string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	fromString, err := pongo2.FromString(str)
	if err != nil {
		return "", err
	}
	return fromString.Execute(data)

}

func ParseContentFile(filename string, data pongo2.Context) (out string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	fromString, err := pongo2.FromCache(filename)
	if err != nil {
		return "", err
	}

	return fromString.Execute(data)

}

func RegisterFunction(name string, fun pongo2.FilterFunction) error {
	if pongo2.FilterExists(name) {
		return pongo2.ReplaceFilter(name, fun)
	}
	return pongo2.RegisterFilter(name, fun)

}

func RegisterFunctionMap(mapf map[string]pongo2.FilterFunction) error {
	for s, function := range mapf {
		err := RegisterFunction(s, function)
		if err != nil {
			return err
		}

	}
	return nil
}
