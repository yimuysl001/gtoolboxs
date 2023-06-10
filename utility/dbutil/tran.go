package dbutil

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func Tran(ctx context.Context, db gdb.DB, f func(ctx2 context.Context, tx gdb.TX)) (err error) {
	begin, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			begin.Commit()
		} else {
			begin.Rollback()
		}
	}()
	err = g.Try(ctx, func(ctx context.Context) {
		f(ctx, begin)
	})
	return
}

func Tran2(ctx context.Context, db gdb.DB, db2 gdb.DB, f func(ctx2 context.Context, tx gdb.TX, tx2 gdb.TX)) (err error) {
	begin, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	begin2, err := db2.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			begin.Commit()
			begin2.Commit()
		} else {
			begin.Rollback()
			begin2.Rollback()
		}
	}()
	err = g.Try(ctx, func(ctx context.Context) {
		f(ctx, begin, begin2)
	})
	return
}
