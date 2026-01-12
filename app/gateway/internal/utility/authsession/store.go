package authsession

import (
	"context"
	redisKey "freeroam/app/gateway/internal/consts/redisKey"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	fieldMemberID   = "member_id"
	fieldVer        = "ver"
	fieldCreatedAt  = "created_at"
	fieldLastSeenAt = "last_seen_at"
	fieldMaxExpAt   = "max_exp_at"
	fieldIP         = "ip"
	fieldUA         = "ua"
	fieldDeviceID   = "device_id"
)

// CreateSession 写入 sess:{sid}（Hash），并设置 TTL
//
// ttlSeconds 必须满足：ttl = min(idle_timeout, max_exp_at - now)
func CreateSession(ctx context.Context, sid string, sess *Session, ttlSeconds int64) error {
	key := redisKey.SessKey(sid)

	if ttlSeconds <= 0 {
		return berror.NewCode(berror.CodeInternal, "ttlSeconds 必须 > 0")
	}

	data := g.Map{
		fieldMemberID:   sess.MemberId,
		fieldVer:        sess.Ver,
		fieldCreatedAt:  sess.CreatedAt,
		fieldLastSeenAt: sess.LastSeenAt,
		fieldMaxExpAt:   sess.MaxExpAt,
	}
	if sess.IP != "" {
		data[fieldIP] = sess.IP
	}
	if sess.UA != "" {
		data[fieldUA] = sess.UA
	}
	if sess.DeviceID != "" {
		data[fieldDeviceID] = sess.DeviceID
	}

	r, err := getRedis(ctx)
	if err != nil {
		return err
	}

	if _, err := r.HSet(ctx, key, data); err != nil {
		return berror.WrapCode(berror.CodeRedisErr, err, "写入会话失败")
	}

	// 设置 TTL（秒）
	if _, err := r.Expire(ctx, key, ttlSeconds); err != nil {
		return berror.WrapCode(berror.CodeRedisErr, err, "设置会话 TTL 失败")
	}
	return nil
}

// GetSession 读取 sess:{sid}，不存在返回 CodeSessionNotExist
func GetSession(ctx context.Context, sid string) (*Session, error) {
	key := redisKey.SessKey(sid)

	r, err := getRedis(ctx)
	if err != nil {
		return nil, err
	}

	v, err := r.HGetAll(ctx, key)
	if err != nil {
		return nil, berror.WrapCode(berror.CodeRedisErr, err, "读取会话失败")
	}
	if v.IsEmpty() {
		return nil, berror.NewCode(berror.CodeTokenInvalid, "会话不存在")
	}

	m := v.Map()
	sess := &Session{
		MemberId:   gconv.Uint64(m[fieldMemberID]),
		Ver:        gconv.Int64(m[fieldVer]),
		CreatedAt:  gconv.Int64(m[fieldCreatedAt]),
		LastSeenAt: gconv.Int64(m[fieldLastSeenAt]),
		MaxExpAt:   gconv.Int64(m[fieldMaxExpAt]),
		IP:         gconv.String(m[fieldIP]),
		UA:         gconv.String(m[fieldUA]),
		DeviceID:   gconv.String(m[fieldDeviceID]),
	}

	// 必要字段基本校验
	if sess.MemberId == 0 || sess.MaxExpAt == 0 {
		return nil, berror.NewCode(berror.CodeTokenInvalid, "会话数据不完整")
	}
	return sess, nil
}

// DeleteSession 删除 sess:{sid}，幂等
func DeleteSession(ctx context.Context, sid string) error {
	key := redisKey.SessKey(sid)
	r, err := getRedis(ctx)
	if err != nil {
		return err
	}
	if _, err := r.Del(ctx, key); err != nil {
		return berror.WrapCode(berror.CodeRedisErr, err, "删除会话失败")
	}
	return nil
}

// ValidateAndTouch 按文档规则校验并续租会话（非原子版本）
//
// 校验项：
// - 会话存在
// - max lifetime：now <= sess.max_exp_at
// - idle timeout：now - sess.last_seen_at <= idle_timeout
// - 一致性：sess.member_id == tokenMemberId 且 sess.ver == tokenVer
//
// 续租规则：
// - ttl = min(idle_timeout, sess.max_exp_at - now)
// - 更新 last_seen_at=now，并 EXPIRE key=ttl
func ValidateAndTouch(
	ctx context.Context,
	sid string,
	tokenMemberId uint64,
	tokenVer int64,
	nowUnix int64,
	idleTimeoutSeconds int64,
) (*Session, int64, error) {
	sess, err := GetSession(ctx, sid)
	if err != nil {
		return nil, 0, err
	}

	// 一致性校验
	if sess.MemberId != tokenMemberId || sess.Ver != tokenVer {
		return nil, 0, berror.NewCode(berror.CodeTokenInvalid, "会话一致性校验失败")
	}

	// 绝对过期（max lifetime）
	if nowUnix > sess.MaxExpAt {
		return nil, 0, berror.NewCode(berror.CodeTokenInvalid, "会话绝对过期")
	}

	// 空闲过期（idle timeout）
	if idleTimeoutSeconds > 0 && nowUnix-sess.LastSeenAt > idleTimeoutSeconds {
		return nil, 0, berror.NewCode(berror.CodeTokenInvalid, "会话空闲过期")
	}

	ttl := CalcTTLSeconds(nowUnix, sess.MaxExpAt, idleTimeoutSeconds)
	if ttl <= 0 {
		return nil, 0, berror.NewCode(berror.CodeTokenInvalid, "会话绝对过期")
	}

	// 更新 last_seen_at + 续租 TTL
	key := redisKey.SessKey(sid)
	r, err := getRedis(ctx)
	if err != nil {
		return nil, 0, err
	}

	if _, err := r.HSet(ctx, key, g.Map{fieldLastSeenAt: nowUnix}); err != nil {
		return nil, 0, berror.WrapCode(berror.CodeRedisErr, err, "更新会话 last_seen_at 失败")
	}
	if _, err := r.Expire(ctx, key, ttl); err != nil {
		return nil, 0, berror.WrapCode(berror.CodeRedisErr, err, "续租会话 TTL 失败")
	}

	// 返回最新 last_seen_at
	sess.LastSeenAt = nowUnix
	return sess, ttl, nil
}
