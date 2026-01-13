package auth

import (
	"context"
	"freeroam/common/model/cjwt"
	"time"

	v1 "freeroam/app/gateway/api/auth/v1"
	"freeroam/app/gateway/internal/service"
	"freeroam/app/gateway/internal/utility/authsession"
	jwtutil "freeroam/app/gateway/internal/utility/jwt"
	orgAuth "freeroam/app/org/api/auth/v1"
	"freeroam/common/berror"
	"freeroam/common/tools/systemConfig"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type sAuth struct {
	AuthRpcService orgAuth.AuthClient
}

func init() {
	conn := grpcx.Client.MustNewGrpcClientConn("org")
	authRpcService := orgAuth.NewAuthClient(conn)
	service.RegisterAuth(&sAuth{authRpcService})
}

func (s *sAuth) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 调用 Org 服务校验账号密码
	if s.AuthRpcService == nil {
		return nil, berror.NewCode(berror.ServiceNotInitialized, "OrgRpcService")
	}
	memberRes, err := s.AuthRpcService.ValidateMemberCredential(ctx, &orgAuth.ValidateMemberCredentialReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 读取系统配置：auth.allow_multi_login（默认不允许多点登录）
	allowMultiLogin := systemConfig.GetBoolD("auth.allow_multi_login", false)
	// max_lifetime：绝对过期（从登录开始最长存活时间）
	maxLifetimeSeconds := systemConfig.GetIntD("auth.session.max_lifetime", int64(86400))
	// idle_timeout：空闲过期（多久不访问则会话失效）
	idleTimeoutSeconds := systemConfig.GetIntD("auth.session.idle_timeout", int64(1800))

	now := time.Now().Unix()
	maxExpAt := now + maxLifetimeSeconds

	// 生成/获取会话版本 ver
	var ver int64
	if allowMultiLogin {
		ver, err = authsession.GetOrInitMemberVersion(ctx, memberRes.MemberId, 1)
	} else {
		ver, err = authsession.IncrMemberVersion(ctx, memberRes.MemberId)
	}
	if err != nil {
		return nil, err
	}

	// 生成 sid（会话ID）-> 使用 JWT jti
	sid := uuid.NewString()

	r := ghttp.RequestFromCtx(ctx)
	sess := &authsession.Session{
		MemberId:   memberRes.MemberId,
		Ver:        ver,
		CreatedAt:  now,
		LastSeenAt: now,
		MaxExpAt:   maxExpAt,
	}
	if r != nil {
		sess.IP = r.GetClientIp()
		sess.UA = r.Header.Get("User-Agent")
	}
	err = authsession.CreateSession(ctx, sid, sess, idleTimeoutSeconds)
	if err != nil {
		return nil, err
	}

	// audience：从请求推断（失败则留空，GenerateToken 会用 jwt.audience 兜底）
	audience := ""
	if r != nil {
		if aud, err := jwtutil.GetAudienceFromRequest(ctx, r); err == nil {
			audience = aud
		}
	}

	// 签发 JWT：member_id/ver 为业务字段；jti(sid)/exp 等为标准字段
	claims := &cjwt.Claims{
		MemberId: memberRes.MemberId,
		Ver:      ver,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        sid, // jti = sid
			ExpiresAt: jwt.NewNumericDate(time.Unix(maxExpAt, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(now, 0)),
		},
	}
	if audience != "" {
		claims.RegisteredClaims.Audience = []string{audience}
	}

	accessToken, err := jwtutil.GenerateToken(ctx, claims)
	if err != nil {
		return nil, berror.WrapCode(berror.CodeInternal, err)
	}

	return &v1.LoginRes{
		AccessToken: accessToken,
		ExpireAt:    maxExpAt,
	}, nil
}

func (s *sAuth) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	// 1. 从请求 Authorization 解析 JWT，得到 sid(jti)
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		// 没有请求上下文时，按幂等成功处理
		return &v1.LogoutRes{Success: true}, nil
	}

	tokenString, err := jwtutil.ExtractTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		// 没 token 也按幂等成功处理
		return &v1.LogoutRes{Success: true}, nil
	}

	claims, err := jwtutil.ValidateToken(ctx, tokenString)
	if err != nil {
		// token 无效/过期也按幂等成功处理（会话可能已过期或已被删除）
		return &v1.LogoutRes{Success: true}, nil
	}
	sid := claims.ID
	if sid != "" {
		// 2. 删除会话：DEL sess:{sid}（幂等）
		_ = authsession.DeleteSession(ctx, sid)
	}

	return &v1.LogoutRes{Success: true}, nil
}
