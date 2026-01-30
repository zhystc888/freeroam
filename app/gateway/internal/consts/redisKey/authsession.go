package redisKey

import "freeroam/common/consts/rediskey"

// 本目录用于集中维护 gateway 服务的 Redis Key 定义（仅本服务使用）。
//
// 约定：
// - 统一根前缀：redisKey.RootPrefix（当前为 "free:"）
// - sid 使用 jti(sid)，类型 string

const (
	// GatewayPrefix 当前服务统一缓存前缀
	GatewayPrefix = rediskey.RootPrefix + "gateway:"

	// SessKeyFmt 会话实体 key：sess:{sid}
	// fmt args: sid
	SessKeyFmt = GatewayPrefix + "sess:%s"
)

func SessKey(sid string) string {
	return rediskey.GetFullKey(SessKeyFmt, sid)
}
