package jwt

import (
	"context"
	"fmt"
	"freeroam/common/model/cjwt"
	"time"

	"freeroam/common/berror"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 签发 JWT Token
// claims: JWT Claims 数据
// 返回: token 字符串和错误
func GenerateToken(ctx context.Context, claims *cjwt.Claims) (string, error) {
	// 读取配置（每次调用都读取，支持 Nacos 动态配置）
	cfg, err := GetConfig(ctx)
	if err != nil {
		return "", berror.WrapCode(berror.CodeConfigReadErr, err)
	}

	// 设置 RegisteredClaims（标准字段只放这里）
	if claims.RegisteredClaims.Issuer == "" {
		claims.RegisteredClaims.Issuer = cfg.Issuer
	}
	if len(claims.RegisteredClaims.Audience) == 0 && cfg.Audience != "" {
		claims.RegisteredClaims.Audience = []string{cfg.Audience}
	}
	if claims.RegisteredClaims.IssuedAt == nil {
		claims.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())
	}
	if claims.RegisteredClaims.ExpiresAt == nil {
		// 兜底：如果调用方没传 max_lifetime 的绝对过期时间，则按配置默认 expire
		// 备注：按你们文档，最终应当使用 max_lifetime 来设置 ExpiresAt（更上层会话逻辑负责）
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Expire) * time.Second))
	}
	if claims.RegisteredClaims.ID == "" {
		return "", berror.NewCode(berror.MissingSid)
	}

	// 创建 token（Claims 已实现 jwt.Claims 接口）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", berror.WrapCode(berror.JWTSigningFailed, err)
	}

	// 注意：如果修改了密钥，之前签发的所有 token 都会无法验证（签名验证失败）
	// 这相当于全体踢下线，需要谨慎操作

	return tokenString, nil
}

// ValidateToken 验证 JWT Token（解析 + 签名验证 + 过期检查）
// tokenString: JWT token 字符串
// ctx: 上下文（用于读取配置）
// 返回: Claims 和错误
func ValidateToken(ctx context.Context, tokenString string) (*cjwt.Claims, error) {
	// 检查 token 是否为空
	if tokenString == "" {
		return nil, berror.NewCode(berror.CodeTokenIsEmpty)
	}

	// 读取配置（每次调用都读取，支持 Nacos 动态配置）
	cfg, err := GetConfig(ctx)
	if err != nil {
		return nil, berror.WrapCode(berror.CodeConfigReadErr, err)
	}

	// 创建空的 Claims 用于解析
	claims := &cjwt.Claims{}

	// 解析并验证 token（包括签名验证）
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, berror.NewCode(berror.CodeTokenInvalid, fmt.Sprintf("不支持的签名算法: %v", token.Header["alg"]))
		}
		// 返回密钥用于签名验证
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		// 其他解析错误
		return nil, berror.NewCode(berror.CodeTokenInvalid, err.Error())
	}

	// 提取 Claims
	if validatedClaims, ok := token.Claims.(*cjwt.Claims); ok && token.Valid {
		return validatedClaims, nil
	}

	return nil, berror.NewCode(berror.TokenClaimsFormatErr)
}

// ExtractTokenFromHeader 从 HTTP Authorization header 中提取 token
// 支持格式: "Bearer <token>" 或直接 "<token>"
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", berror.NewCode(berror.CodeTokenIsEmpty)
	}

	// 移除 "Bearer " 前缀（如果存在）
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):], nil
	}

	// 如果没有 Bearer 前缀，直接返回（兼容直接传 token 的情况）
	return authHeader, nil
}
