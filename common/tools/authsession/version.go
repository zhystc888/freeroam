package authsession

import (
	"context"
	"freeroam/common/berror"
	"freeroam/common/consts/rediskey"
)

// GetMemberVersion 获取成员会话版本号（auth:ver:member:{memberId}）
// 如果不存在则返回 0（注意：与 GetOrInitMemberVersion 不同，本方法不会初始化）
func GetMemberVersion(ctx context.Context, memberId uint64) (int64, error) {
	key := rediskey.MemberVerKey(memberId)

	r, err := GetRedis(ctx)
	if err != nil {
		return 0, err
	}

	v, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 member ver 失败")
	}
	if v.IsEmpty() {
		return 0, nil
	}
	return v.Int64(), nil
}

// GetOrInitMemberVersion 获取成员会话版本号（auth:ver:member:{memberId}）
// 如果不存在则初始化为 initVal（通常为 1），并返回当前版本
func GetOrInitMemberVersion(ctx context.Context, memberId uint64, initVal int64) (int64, error) {
	key := rediskey.MemberVerKey(memberId)

	r, err := GetRedis(ctx)
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

// GetGlobalVersion 获取全局会话版本号（auth:ver:global）
// 如果不存在则返回 0（注意：与 GetOrInitGlobalVersion 不同，本方法不会初始化）
func GetGlobalVersion(ctx context.Context) (int64, error) {
	key := rediskey.GlobalVerKey()

	r, err := GetRedis(ctx)
	if err != nil {
		return 0, err
	}

	v, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 global ver 失败")
	}
	if v.IsEmpty() {
		return 0, nil
	}
	return v.Int64(), nil
}

// GetOrInitGlobalVersion 获取全局会话版本号（auth:ver:global）
// 如果不存在则初始化为 initVal（通常为 1），并返回当前版本
func GetOrInitGlobalVersion(ctx context.Context, initVal int64) (int64, error) {
	key := rediskey.GlobalVerKey()

	r, err := GetRedis(ctx)
	if err != nil {
		return 0, err
	}

	v, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 global ver 失败")
	}
	if !v.IsEmpty() {
		return v.Int64(), nil
	}

	// SETNX key initVal
	if _, err := r.SetNX(ctx, key, initVal); err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "初始化 global ver 失败")
	}

	v2, err := r.Get(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "读取 global ver 失败")
	}
	if v2.IsEmpty() {
		return initVal, nil
	}
	return v2.Int64(), nil
}

// IncrMemberVersion INCR 成员会话版本号，并返回递增后的版本
func IncrMemberVersion(ctx context.Context, memberId uint64) (int64, error) {
	key := rediskey.MemberVerKey(memberId)
	r, err := GetRedis(ctx)
	if err != nil {
		return 0, err
	}
	v, err := r.Incr(ctx, key)
	if err != nil {
		return 0, berror.WrapCode(berror.CodeRedisErr, err, "INCR member ver 失败")
	}
	return v, nil
}
