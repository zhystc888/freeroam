package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT Claims 结构体定义
// 根据文档：成员登录与会话管理方案
type Claims struct {
	// 业务字段（仅业务相关，避免与 JWT 标准字段冲突）
	MemberId uint64 `json:"member_id"` // 成员ID
	Ver      int64  `json:"ver"`       // 成员会话版本

	// JWT 标准字段：只放在 RegisteredClaims（避免 exp/iat/iss/aud/jti 重复字段导致歧义）
	// - sid 使用 jti：RegisteredClaims.ID
	// - exp 使用 ExpiresAt；iat 使用 IssuedAt；iss 使用 Issuer；aud 使用 Audience
	jwt.RegisteredClaims
}
