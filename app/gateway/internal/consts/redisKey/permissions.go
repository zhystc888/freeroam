package redisKey

import (
	"freeroam/common/consts/rediskey"
)

// 本目录用于集中维护 org 服务的 Redis Key 定义（仅本服务使用）。
//
// 约定：
// - 统一根前缀：redisKey.RootPrefix（当前为 "free:"）

const (
	// PermissionsKeyFmt 会话实体 key：permissions:{permTypes}:{memberId}:{globalVer}:{ver}
	PermissionsKeyFmt = GatewayPrefix + "permissions:%s:%d:%d:%d"
)

func PermissionsKey(permType string, memberId uint64, globalVer, ver int64) string {
	return rediskey.GetFullKey(PermissionsKeyFmt, permType, memberId, globalVer, ver)
}
