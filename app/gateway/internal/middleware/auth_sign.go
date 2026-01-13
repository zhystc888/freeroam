package middleware

import (
	"freeroam/app/gateway/internal/utility/authsession"
	"freeroam/common/tools/systemConfig"
	"time"

	jwtutil "freeroam/app/gateway/internal/utility/jwt"
	"freeroam/common/berror"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const CtxKeyJwtClaims = "auth.jwt.claims"

// AuthSign JWT 验签中间件
func AuthSign(r *ghttp.Request) {
	tokenString, err := jwtutil.ExtractTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		writeAuthError(r, err)
		return
	}

	claims, err := jwtutil.ValidateToken(r.Context(), tokenString)
	if err != nil {
		writeAuthError(r, err)
		return
	}

	// 校验会话版本 & 空闲过期，并续租
	idleTimeoutSeconds := systemConfig.GetIntD("auth.session.idle_timeout", int64(1800))
	
	if _, _, err := authsession.ValidateAndTouch(
		r.Context(),
		claims.ID,       // sid = jti
		claims.MemberId, // token 内的 member_id
		claims.Ver,      // token 内的 ver
		time.Now().Unix(),
		idleTimeoutSeconds,
	); err != nil {
		writeAuthError(r, err)
		return
	}

	ctx := grpcx.Ctx.SetOutgoing(r.Context(), g.Map{
		CtxKeyJwtClaims: claims,
	})
	r.SetCtx(ctx)

	r.Middleware.Next()
}

func writeAuthError(r *ghttp.Request, err error) {
	if gerror.Code(err) == nil {
		err = berror.NewCode(berror.CodeTokenInvalid, err.Error())
	}
	r.SetError(err)
}
