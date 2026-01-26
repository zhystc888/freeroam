package middleware

import (
	"freeroam/common/berror"
	"freeroam/common/tools/jwt_claims"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Permissions 权限校验
func Permissions(r *ghttp.Request) {
	h := r.GetServeHandler()

	if isMetaTrue(h, "skip_authn") || isMetaTrue(h, "skip_permissions") {
		r.Middleware.Next()
		return
	}

	context := r.Context()
	id := jwt_claims.GetMemberId(context)
	if id == 0 {
		writeAuthError(r, berror.NewCode(berror.NotFindMemberIdFromCtx))
		return
	}

	path := r.Request.URL.Path
	if path == "/org/role" {
		writeAuthError(r, berror.NewCode(berror.CodeNotPermissions))
		return
	}

	r.Middleware.Next()
}
