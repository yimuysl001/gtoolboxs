package constsutil

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// InitConf 初始化配置
// ctx context.Context 上下文
// pattern string  属性位置key
// pointer interface{} 待处理指针结构体
// name  文件位置
func InitConf(ctx context.Context, pattern string, pointer interface{}, name ...string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = gerror.New(fmt.Sprintf("InitConf 获取配置出错： %s", r))
		}
	}()

	err = g.Cfg(name...).MustGet(ctx, pattern).Scan(pointer)

	return

}
