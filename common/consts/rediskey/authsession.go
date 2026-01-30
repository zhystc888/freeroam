package rediskey

// 本目录用于集中维护 gateway 服务的 Redis Key 定义（仅本服务使用）。
//
// 约定：
// - 统一根前缀：redisKey.RootPrefix（当前为 "free:"）
// - sid 使用 jti(sid)，类型 string

const (
	// GatewayPrefix 当前服务统一缓存前缀
	GatewayPrefix = RootPrefix + "gateway:"

	// MemberVerKeyFmt 成员会话版本 key：auth:ver:member:{memberId}
	// fmt args: memberId
	MemberVerKeyFmt = GatewayPrefix + "auth:ver:member:%d"

	// GlobalVerKeyFmt 全局会话版本 key：auth:ver:global
	GlobalVerKeyFmt = GatewayPrefix + "auth:ver:global"
)

func MemberVerKey(memberId uint64) string {
	return GetFullKey(MemberVerKeyFmt, memberId)
}

func GlobalVerKey() string {
	return GlobalVerKeyFmt
}
