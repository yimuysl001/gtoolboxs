package v0

import (
	"context"
	"embed"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yimuysl001/gtoolboxs/utility/constsutil"
	"github.com/yimuysl001/gtoolboxs/utility/fileutil"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"github.com/yimuysl001/gtoolboxs/utility/myctx"
	"os"
)

//go:embed knife
var Files embed.FS

var serverRoot string

func init() {
	ctx := myctx.New()
	constsutil.InitConf(ctx, "server.serverRoot", &serverRoot)

	if serverRoot != "" {
		serverRoot = serverRoot + "/"
		InitFile(gctx.New(), "knife")
	}

}

func InitFile(ctx context.Context, name string) {
	exists, _ := fileutil.PathExists(serverRoot + name)
	if !exists {
		if err := os.MkdirAll(serverRoot+name, os.ModePerm); err != nil {
			logger.Logger.ErrorCtx(ctx, err)
			return
		}
	}

	file, err := Files.ReadDir(name)
	if err != nil {
		logger.Logger.ErrorCtx(ctx, err)
	}
	for _, f := range file {
		info, _ := f.Info()
		namepath := name + "/" + info.Name()
		if f.IsDir() {
			InitFile(ctx, namepath)
			continue
		}
		exists, _ = fileutil.PathExists(serverRoot + namepath)
		if exists {
			continue
		}
		readFile, err := Files.ReadFile(namepath)
		if err != nil {
			logger.Logger.ErrorCtx(ctx, err)
			continue
		}
		err = os.WriteFile(serverRoot+namepath, readFile, info.Mode())
		logger.Logger.IfErrorCtx(ctx, err)

	}

}
