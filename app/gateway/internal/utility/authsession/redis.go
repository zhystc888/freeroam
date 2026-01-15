package authsession

import (
	"context"

	"freeroam/common/berror"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// getRedis 获取 Redis 实例（判空）。
//
// g.Redis() 在以下场景可能为 nil：
// - 未引入 redis driver（_ "github.com/gogf/gf/contrib/nosql/redis/v2"）
// - 未配置 redis.default（或配置加载失败）
func getRedis(ctx context.Context) (*gredis.Redis, error) {
	r := g.Redis()
	if r == nil {
		return nil, berror.NewCode(berror.RedisNotInitialized)
	}
	return r, nil
}
