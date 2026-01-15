package jwt

import (
	"context"

	"freeroam/common/berror"

	"github.com/gogf/gf/v2/frame/g"
)

// Config JWT 配置结构体
type Config struct {
	Secret   string // JWT 密钥（必须）
	Issuer   string // 签发者（可选，用于标识 token 的签发者，如 "freeroam-gateway"）
	Audience string // 默认接收者（可选，用于标识 token 的目标接收者，如 "freeroam-web-app"、"freeroam-mobile-app"）
	Expire   int64  // 默认过期时间（秒，可选，实际使用 max_lifetime）
}

// GetConfig 从配置文件读取 JWT 配置
// 配置来源：
//   - 本地环境（APP_ENV=local）: app/gateway/manifest/config/config.yaml
//   - 生产环境: Nacos 配置中心的 gateway-config.yaml (DataId: gateway-config.yaml, Group: APP_ENV)
//   - 注意：密钥不应存储在数据库中（安全考虑）
//
// 重要提示：
//   - 修改密钥会导致所有已签发的 token 无法验证（签名验证失败），相当于全体踢下线
//   - 如需更换密钥，建议先通知用户，或实现密钥版本管理（本期不实现）
//   - 配置支持动态读取（Nacos 配置中心支持动态更新，但需重启服务才能生效）
func GetConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{
		Secret:   g.Cfg().MustGet(ctx, "jwt.secret", "freeroam-gateway-default-secret-key-change-in-production").String(),
		Issuer:   g.Cfg().MustGet(ctx, "jwt.issuer", "freeroam-gateway").String(),
		Audience: g.Cfg().MustGet(ctx, "jwt.audience", "freeroam-client").String(), // 默认接收者，可在登录时覆盖
		Expire:   g.Cfg().MustGet(ctx, "jwt.expire", 86400).Int64(),                // 默认 24 小时（兜底，实际以 sess.max_exp_at 为准）
	}

	// 如果密钥为空或使用默认值，在非生产环境允许（开发环境）
	// 生产环境应该强制配置密钥
	if cfg.Secret == "" {
		return nil, berror.NewCode(berror.JWTSecretCannotBeEmpty)
	}

	return cfg, nil
}
