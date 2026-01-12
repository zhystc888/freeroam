package authsession

import (
	"context"

	redisKey "freeroam/app/gateway/internal/consts/redisKey"
	"freeroam/common/berror"
)

// GetOrInitMemberVersion 获取成员会话版本号（auth:ver:member:{memberId}）
// 如果不存在则初始化为 initVal（通常为 1），并返回当前版本
func GetOrInitMemberVersion(ctx context.Context, memberId uint64, initVal int64) (int64, error) {
	key := redisKey.MemberVerKey(memberId)

	r, err := getRedis(ctx)
	if err != nil {
		return 0, err
	}

	v, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 member ver 失败")
	}
	if !v.IsEmpty() {
		return v.Int64(), nil
	}

	// SETNX key initVal
	if _, err := r.SetNX(ctx, key, initVal); err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "初始化 member ver 失败")
	}

	v2, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 member ver 失败")
	}
	if v2.IsEmpty() {
		return initVal, nil
	}
	return v2.Int64(), nil
}

// IncrMemberVersion INCR 成员会话版本号，并返回递增后的版本
func IncrMemberVersion(ctx context.Context, memberId uint64) (int64, error) {
	key := redisKey.MemberVerKey(memberId)
	r, err := getRedis(ctx)
	if err != nil {
		return 0, err
	}
	v, err := r.Incr(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "INCR member ver 失败")
	}
	return v, nil
}
