package timeutil

import (
	"context"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"time"
)

const Farmat = "2006-01-02"
const Farmat1 = "2006-01-02 15:04:05"
const Farmat2 = "2006-01-02 15:04:05.999"

func StrToTime(str string) time.Time {
	var err error
	parse := time.Time{}
	switch len(str) {
	case len(Farmat):
		parse, err = time.ParseInLocation(Farmat, str, time.Local)
		logger.Logger.PanicErrorCtx(context.Background(), err)
	case len(Farmat1):
		parse, err = time.ParseInLocation(Farmat1, str, time.Local)
		logger.Logger.PanicErrorCtx(context.Background(), err)
	case len(Farmat2):
		parse, err = time.ParseInLocation(Farmat2, str, time.Local)
		logger.Logger.PanicErrorCtx(context.Background(), err)
	}

	return parse

}
