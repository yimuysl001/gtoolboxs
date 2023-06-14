package sys2

import (
	"github.com/yimuysl001/gtoolboxs/utility/sys"
)

func init() {
	//校验同名程序是否启动
	sys.CheckSys()
	//守护线程
	go sys.BlockingThread()
}
