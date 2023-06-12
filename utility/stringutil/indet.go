package stringutil

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/beevik/etree"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

func IndentJson(ctx context.Context, str string) string {
	var bb bytes.Buffer
	logger.Logger.PanicErrorCtx(ctx, json.Indent(&bb, []byte(str), "", "    "))
	return bb.String()

}

func IndentXml(ctx context.Context, str string) string {

	document := etree.NewDocument()
	logger.Logger.PanicErrorCtx(ctx, document.ReadFromString(str))
	document.IndentTabs()
	toString, err := document.WriteToString()
	logger.Logger.PanicErrorCtx(ctx, err)
	return toString

}
