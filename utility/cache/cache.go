package cache

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/yimuysl001/gtoolboxs/utility/cache/file"
	leveldbcache "github.com/yimuysl001/gtoolboxs/utility/cache/leveldb"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
)

type CacheType string

const (
	Redis   CacheType = "redis"
	File    CacheType = "file"
	LevelDb CacheType = "levleldb"
)

// cache 缓存驱动
var cache *gcache.Cache
var adapter gcache.Adapter

func GetAdapter() gcache.Adapter {
	return adapter
}

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		SetAdapter(context.Background(), "", "")
		//panic("cache uninitialized.")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context, t CacheType, path string) {
	//var adapter gcache.Adapter
	switch t {
	case Redis:
		adapter = gcache.NewAdapterRedis(g.Redis())
	case LevelDb:
		if path == "" {
			logger.Logger.ErrorCtx(ctx, "file path must be configured for file caching.")
			return
		}

		if !gfile.Exists(path) {
			if err := gfile.Mkdir(path); err != nil {
				g.Log().Fatalf(ctx, "Failed to create the cache directory. Procedure, err:%+v", err)
				return
			}
		}
		adapter = leveldbcache.NewAdapterFile(path)
	case File:
		if path == "" {
			logger.Logger.ErrorCtx(ctx, "file path must be configured for file caching.")
			return
		}

		if !gfile.Exists(path) {
			if err := gfile.Mkdir(path); err != nil {
				g.Log().Fatalf(ctx, "Failed to create the cache directory. Procedure, err:%+v", err)
				return
			}
		}
		adapter = file.NewAdapterFile(path)
	default:
		adapter = gcache.NewAdapterMemory()
	}

	g.DB().GetCache().SetAdapter(adapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(adapter)

}
