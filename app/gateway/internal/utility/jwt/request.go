package jwt

import (
	"context"
	"strings"

	"freeroam/common/berror"

	"github.com/gogf/gf/v2/net/ghttp"
)

// GetAudienceFromRequest 从 HTTP 请求中获取 audience（接收者）
// 获取优先级：
//  1. 请求参数 client_type（如果存在）
//  2. 请求头 X-Client-Type（如果存在）
//  3. 从 User-Agent 推断客户端类型（如果包含特定关键字）
//  4. 配置文件中的默认值（jwt.audience）
//
// 客户端类型映射：
//   - "web" -> "freeroam-web-app"
//   - "mobile" / "app" / "ios" / "android" -> "freeroam-mobile-app"
//   - 其他 -> 保持原值或使用默认值
func GetAudienceFromRequest(ctx context.Context, r *ghttp.Request) (string, error) {
	// 1. 优先从请求参数获取 client_type
	clientType := r.Get("client_type").String()
	if clientType != "" {
		return mapClientTypeToAudience(clientType)
	}

	// 2. 从请求头 X-Client-Type 获取
	clientType = r.Header.Get("X-Client-Type")
	if clientType != "" {
		return mapClientTypeToAudience(clientType)
	}

	// 3. 从 User-Agent 推断（可选，可根据业务需求调整）
	userAgent := r.Header.Get("User-Agent")
	if userAgent != "" {
		ua := strings.ToLower(userAgent)
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "android") || strings.Contains(ua, "ios") || strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
			return "freeroam-mobile-app", nil
		}
		if strings.Contains(ua, "mozilla") || strings.Contains(ua, "chrome") || strings.Contains(ua, "firefox") || strings.Contains(ua, "safari") {
			return "freeroam-web-app", nil
		}
	}

	// 4. 使用配置中的默认值
	cfg, err := GetConfig(ctx)
	if err != nil {
		return "", err
	}
	return cfg.Audience, nil
}

// mapClientTypeToAudience 将客户端类型映射为 audience
func mapClientTypeToAudience(clientType string) (string, error) {
	clientType = strings.ToLower(strings.TrimSpace(clientType))

	switch clientType {
	case "web", "browser", "pc":
		return "freeroam-web-app", nil
	case "mobile", "app", "ios", "android", "iphone", "ipad":
		return "freeroam-mobile-app", nil
	default:
		// 如果传入了自定义的 client_type，直接使用（或添加前缀）
		// 这里可以根据业务需求调整
		if clientType != "" {
			return "freeroam-" + clientType + "-app", nil
		}
		// 如果没有匹配，返回错误（需要在调用处处理）
		return "", berror.NewCode(berror.UnrecognizedClientType, clientType)
	}
}
