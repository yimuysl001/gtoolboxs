package hello

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/yimuysl001/gtoolboxs/api/hello/v1"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	one, err := g.DB().GetOne(ctx, "select * from test")
	fmt.Println(one)
	fmt.Println(err)

	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
