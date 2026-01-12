package authsession

import "time"

// CalcTTLSeconds 根据文档公式计算续租 TTL（秒）
//
// ttl = min(idle_timeout, max_exp_at - now)
//
// 返回：
// - ttlSeconds: 续租 TTL（秒），<=0 表示绝对过期
func CalcTTLSeconds(nowUnix int64, maxExpAtUnix int64, idleTimeoutSeconds int64) int64 {
	remain := maxExpAtUnix - nowUnix
	if remain <= 0 {
		return 0
	}
	if idleTimeoutSeconds <= 0 {
		// 没配 idle_timeout 时，不做滑动过期（仅受 max lifetime 控制）
		return remain
	}
	if remain < idleTimeoutSeconds {
		return remain
	}
	return idleTimeoutSeconds
}

func Seconds(d time.Duration) int64 {
	return int64(d / time.Second)
}
