package redisutil

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

func GetRedis(name ...string) *gredis.Redis {
	return g.Redis(name...)

}
