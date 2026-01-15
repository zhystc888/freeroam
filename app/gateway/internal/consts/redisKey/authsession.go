package redisKey

import "freeroam/common/consts/rediskey"

// 本目录用于集中维护 gateway 服务的 Redis Key 定义（仅本服务使用）。
//
// 约定：
// - 统一根前缀：rediskey.RootPrefix（当前为 "free"）
// - sid 使用 jti(sid)，类型 string

const (
	// SessKeyFmt 会话实体 key：sess:{sid}
	// fmt args: sid
	SessKeyFmt = rediskey.RootPrefix + ":gateway:sess:%s"

	// MemberVerKeyFmt 成员会话版本 key：auth:ver:member:{memberId}
	// fmt args: memberId
	MemberVerKeyFmt = rediskey.RootPrefix + ":gateway:auth:ver:member:%d"
)

func SessKey(sid string) string {
	return rediskey.GetFullKey(SessKeyFmt, sid)
}

func MemberVerKey(memberId uint64) string {
	return rediskey.GetFullKey(MemberVerKeyFmt, memberId)
}
