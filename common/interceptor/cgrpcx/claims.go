package cgrpcx

import (
	"context"
	"freeroam/common/consts"
	"freeroam/common/model/cjwt"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

// UnaryServerInterceptorInjectJwtClaims 是一个 gRPC Unary Server 拦截器，用于：
// 1) 将 incoming metadata 复制到 outgoing metadata（方便后续下游调用透传）。
// 2) 将 incoming 里解析出的 JWT Claims 写入 ctx 的 value（便于业务逻辑层直接从 ctx 取）。
//
// 注意：context 是不可变对象，必须把 context.WithValue 返回的新 ctx 继续向下传递给 handler 才会生效。
func UnaryServerInterceptorInjectJwtClaims(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	grpcx.Ctx.IncomingToOutgoing(ctx)

	ctx = withJwtClaimsFromIncoming(ctx)
	return handler(ctx, req)
}

// withJwtClaimsFromIncoming 从 grpcx 的 incoming 上下文映射中读取 claims，并返回注入了 claims 的新 ctx。
// 如果不存在/类型不匹配，则原样返回 ctx。
func withJwtClaimsFromIncoming(ctx context.Context) context.Context {
	incomingMap := grpcx.Ctx.IncomingMap(ctx)
	rawClaims := incomingMap.Get(consts.CtxKeyJwtClaims)
	if rawClaims == nil {
		return ctx
	}

	claims, ok := rawClaims.(*cjwt.Claims)
	if !ok || claims == nil {
		return ctx
	}

	return context.WithValue(ctx, consts.CtxKeyJwtClaims, claims)
}
