package rediskey

import (
	"fmt"
)

const (
	// RootPrefix Redis key 根前缀（固定，不做配置化拼接）
	RootPrefix = "free"
)

// GetFullKey 将 keyFmt 和 args 通过 fmt.Sprintf 生成最终 Redis key。
//
// 约定：在 keyFmt 定义时就写好占位符，例如：
//
//	const SessKeyFmt = "free:gateway:sess:%s" // %s: sid
//	key := rediskey.GetFullKey(SessKeyFmt, sid)
func GetFullKey(keyFmt string, args ...any) string {
	return fmt.Sprintf(keyFmt, args...)
}
