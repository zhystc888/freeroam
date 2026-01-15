package cgrpcx

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
)

func ErrorLogInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	res, err := handler(ctx, req)

	if err == nil {
		return res, nil
	}

	code := gerror.Code(err)
	g.Log().Warningf(ctx, "grpc error: %v", err)
	err = gerror.NewCode(code, code.Message())
	return res, err
}
