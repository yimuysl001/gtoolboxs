package main

import (
	"fmt"
	_ "github.com/yimuysl001/gtoolboxs/internal/packed"
	"github.com/yimuysl001/gtoolboxs/utility/sys/autolnk"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())

	//all, err := dbutil.DB("log").Model("TBZDBQ").All()
	//logger.Logger.Info(err)
	//logger.Logger.Info(all)
	//
	//all, err = dbutil.DB("log").Model("TBZDBQ").All()
	//logger.Logger.Info(err)
	//logger.Logger.Info(all)

	//m := gmap.New()
	//
	//m.Set("id", 4)
	//m.Set("name", "de2")
	//
	//dbutil.DB().Model("test").Insert(m)
	fmt.Println(autolnk.CreateDesktop())
	fmt.Println(autolnk.CreateStartup())

}
