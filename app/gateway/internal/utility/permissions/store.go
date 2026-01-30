package permissions

import (
	"context"
	"freeroam/app/gateway/internal/consts/enum"
	"freeroam/app/gateway/internal/consts/redisKey"
	"freeroam/common/berror"
	"freeroam/common/tools/authsession"
	"freeroam/common/tools/jwt_claims"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
)

const (
	ttlSeconds = 1800 // 权限缓存时常 30*60 秒
)

// 获取登录用户权限缓存 key
func getLoginMemberRedisKey(ctx context.Context, permType string) (string, error) {
	// 获取当前登录用户信息
	memberId := jwt_claims.GetMemberId(ctx)
	globalVer := jwt_claims.GetGlobalVer(ctx)
	ver := jwt_claims.GetMemberVer(ctx)
	if memberId == 0 || globalVer == 0 || ver == 0 {
		return "", gerror.NewCodef(berror.CodeTokenIsEmpty,
			"获取登录用户数据失败, memberId:%d, globalVer:%d, ver:%d", memberId, globalVer, ver)
	}

	// 获取缓存 key
	return redisKey.PermissionsKey(permType, memberId, globalVer, ver), nil
}

// GetFrontPermissions 获取前端缓存中权限
// 数据不存在 返回 nil, nil
func GetFrontPermissions(ctx context.Context) ([]string, error) {
	key, err := getLoginMemberRedisKey(ctx, enum.PermissionsFront)
	if err != nil {
		return nil, err
	}

	redis, err := authsession.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	rdata, err := redis.Get(ctx, key)
	if err != nil {
		return nil, gerror.WrapCode(berror.RedisErr, err, "获取权限缓存失败")
	}
	if rdata.IsNil() {
		return nil, nil
	}

	// 续租
	if _, err = redis.Expire(ctx, key, ttlSeconds); err != nil {
		logger.Errorf("续租会话 TTL 失败, err: %v", err)
	}
	return rdata.Strings(), nil
}

// SaveFrontPermissions 保存前端权限到缓存中
func SaveFrontPermissions(ctx context.Context, permissions []string) error {
	key, err := getLoginMemberRedisKey(ctx, enum.PermissionsFront)
	if err != nil {
		return err
	}

	redis, err := authsession.GetRedis(ctx)
	if err != nil {
		return err
	}

	// 放入缓存
	_, err = redis.Set(ctx, key, permissions)
	if err != nil {
		return gerror.WrapCode(berror.RedisErr, err, "设置权限缓存失败")
	}

	// 设置 TTL（秒）
	if _, err = redis.Expire(ctx, key, ttlSeconds); err != nil {
		return gerror.WrapCode(berror.CodeRedisErr, err, "设置会话 TTL 失败")
	}

	return nil
}

// VeryApiPermissions 验证 api 权限
// 1:key不存在 2:权限校验成功 3:权限校验失
func VeryApiPermissions(ctx context.Context, permissions string) (int, error) {
	key, err := getLoginMemberRedisKey(ctx, enum.PermissionsApi)
	if err != nil {
		return 0, err
	}

	redis, err := authsession.GetRedis(ctx)
	if err != nil {
		return 0, err
	}

	// 缓存 key 是否存在
	exists, err := redis.Exists(ctx, key)
	if err != nil {
		return 0, gerror.WrapCode(berror.RedisErr, err, "获取权限缓存失败")
	}
	if exists == 0 {
		return 1, nil
	}

	// 续租
	if _, err = redis.Expire(ctx, key, ttlSeconds); err != nil {
		logger.Errorf("续租会话 TTL 失败, err: %v", err)
	}

	// 缓存中是否存在权限
	member, err := redis.SIsMember(ctx, key, permissions)
	if err != nil {
		return 0, gerror.WrapCode(berror.RedisErr, err, "验证权限失败")
	}
	if member == 0 {
		return 3, nil
	}

	return 2, nil
}

// SaveApiPermissions 保存 api 限到缓存中
func SaveApiPermissions(ctx context.Context, permissions []string) error {
	key, err := getLoginMemberRedisKey(ctx, enum.PermissionsApi)
	if err != nil {
		return err
	}

	redis, err := authsession.GetRedis(ctx)
	if err != nil {
		return err
	}

	// 放入缓存
	if len(permissions) == 0 {
		return nil
	}
	permArgs := make([]interface{}, len(permissions))
	for i, v := range permissions {
		permArgs[i] = v
	}
	_, err = redis.SAdd(ctx, key, permArgs[0], permArgs[1:]...)
	if err != nil {
		return gerror.WrapCode(berror.RedisErr, err, "设置权限缓存失败")
	}

	// 设置 TTL（秒）
	if _, err = redis.Expire(ctx, key, ttlSeconds); err != nil {
		return gerror.WrapCode(berror.CodeRedisErr, err, "设置会话 TTL 失败")
	}

	return nil
}
