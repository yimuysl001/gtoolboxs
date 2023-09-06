package myctx

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

type Context interface {
	context.Context
	Cancel()
}

type MyContext struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (m *MyContext) Deadline() (deadline time.Time, ok bool) {
	return m.ctx.Deadline()
}

func (m *MyContext) Done() <-chan struct{} {
	return m.ctx.Done()
}

func (m *MyContext) Err() error {
	return m.ctx.Err()
}

func (m *MyContext) Value(key any) any {
	return m.ctx.Value(key)
}

func (m *MyContext) Cancel() {
	m.cancel()
}

func New() *MyContext {
	ctx, c := context.WithCancel(gctx.New())
	return &MyContext{
		ctx:    ctx,
		cancel: c,
	}
}

func DoLoopCtx(f func(ctx context.Context), cywlx ...string) (Context, error) {
	ctx := New()
	_, err := DoLoop(ctx, f, cywlx...)

	return ctx, err
}

func DoLoop(ctx Context, f func(ctx context.Context), cywlx ...string) (*gcron.Entry, error) {
	return gcron.AddOnce(ctx, "@every 1s", func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				f(ctx)
			}

		}
	}, cywlx...)

}
