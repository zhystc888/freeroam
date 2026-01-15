package jwt_claims

import (
	"context"
	"freeroam/common/consts"
	"freeroam/common/model/cjwt"
	"time"
)

// GetClaims 从 ctx 中获取业务定义的 JWT Claims（`*cjwt.Claims`）。
//
// 约定：claims 由鉴权中间件/拦截器写入 ctx，key 为 `consts.CtxKeyJwtClaims`。
// 如果 ctx 中不存在 claims，或类型不匹配，则返回 nil。
func GetClaims(ctx context.Context) *cjwt.Claims {
	if ctx == nil {
		return nil
	}
	rawClaims := ctx.Value(consts.CtxKeyJwtClaims)
	if rawClaims == nil {
		return nil
	}

	claims, ok := rawClaims.(*cjwt.Claims)
	if !ok || claims == nil {
		return nil
	}

	return claims
}

// GetMemberId 从 ctx 的 claims 中获取 MemberId。
// 当 claims 不存在时返回 0。
func GetMemberId(ctx context.Context) uint64 {
	claims := GetClaims(ctx)
	if claims == nil {
		return 0
	}
	return claims.MemberId
}

// GetMemberVer 从 ctx 的 claims 中获取会话版本 Ver。
// 当 claims 不存在时返回 0。
func GetMemberVer(ctx context.Context) int64 {
	claims := GetClaims(ctx)
	if claims == nil {
		return 0
	}
	return claims.Ver
}

// GetSid 从 ctx 的 claims 中获取 sid（约定使用 JWT 标准字段 jti：`RegisteredClaims.ID`）。
// 当 claims 不存在时返回空字符串。
func GetSid(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	return claims.ID
}

// GetExp 从 ctx 的 claims 中获取过期时间 exp。
// 当 claims 不存在或未设置 exp 时返回 nil。
func GetExp(ctx context.Context) *time.Time {
	claims := GetClaims(ctx)
	if claims == nil {
		return nil
	}
	if claims.ExpiresAt == nil {
		return nil
	}
	return &claims.ExpiresAt.Time
}
