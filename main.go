package main

import (
	_ "github.com/yimuysl001/gtoolboxs/internal/packed"
	"github.com/yimuysl001/gtoolboxs/utility/dbutil"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())

	all, err := dbutil.DB("TEST").Model("test").All()
	logger.Logger.Info(err)
	logger.Logger.Info(all)

	all, err = dbutil.DB("TEST").Model("test").All()
	logger.Logger.Info(err)
	logger.Logger.Info(all)

	//m := gmap.New()
	//
	//m.Set("id", 4)
	//m.Set("name", "de2")
	//
	//dbutil.DB().Model("test").Insert(m)

}
