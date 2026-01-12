package authsession

// Session 对应 Redis 中 sess:{sid} 的会话实体（Hash）
//
// 字段来源：成员登录与会话管理方案（YWNYwGVBjitRswklGETcSI4NnWb）
// 必要字段：
// - member_id
// - ver
// - created_at
// - last_seen_at
// - max_exp_at
//
// 可选字段（审计/定位用，可按需写入）：
// - ip
// - ua
// - device_id
type Session struct {
	MemberId   uint64
	Ver        int64
	CreatedAt  int64
	LastSeenAt int64
	MaxExpAt   int64

	IP       string
	UA       string
	DeviceID string
}
