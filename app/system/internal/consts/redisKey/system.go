package redisKey

import "freeroam/common/consts/rediskey"

// 本目录用于集中维护 system 服务的 Redis Key 定义（仅本服务使用）。
//
// 约定：
// - 统一根前缀：redisKey.RootPrefix（当前为 "free:"）

const (
	// SystemPrefix 当前服务统一缓存前缀
	SystemPrefix = rediskey.RootPrefix + "system:"

	// EnumKeyFmt 枚举缓存前缀 key：enum:{enumType}
	EnumKeyFmt = SystemPrefix + "enum:%s"

	// SystemConfigKeyFmt 系统配置缓存前缀 key：config:
	SystemConfigKeyFmt = SystemPrefix + "config:"
)

func EnumKey(enumType string) string {
	return rediskey.GetFullKey(EnumKeyFmt, enumType)
}

func SystemConfigKey() string {
	return SystemConfigKeyFmt
}
