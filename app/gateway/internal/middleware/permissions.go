package middleware

import (
	"freeroam/app/gateway/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Permissions 权限校验
func Permissions(r *ghttp.Request) {
	h := r.GetServeHandler()

	// 配置了无需登录，不校验权限
	if isMetaTrue(h, "skip_authn") {
		r.Middleware.Next()
		return
	}

	tagPermissions := h.GetMetaTag("perm")
	// 没有获取到权限标记，直接跳过验证
	if tagPermissions == "" {
		r.Middleware.Next()
		return
	}

	// 校验权限
	apiPermissions, err := service.Permissions().VeryApiPermissions(r.Context(), tagPermissions)
	if err != nil {
		writeAuthError(r, err)
		return
	}

	// 无权限
	if !apiPermissions {
		writeAuthError(r, gerror.NewCodef(berror.CodeNotPermissions, "没有 %s 权限", tagPermissions))
		return
	}

	r.Middleware.Next()
}
