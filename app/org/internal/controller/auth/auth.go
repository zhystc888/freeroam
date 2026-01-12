package auth

import (
	"context"
	v1 "freeroam/app/org/api/auth/v1"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAuthServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAuthServer(s.Server, &Controller{})
}

func (*Controller) ValidateMemberCredential(ctx context.Context, req *v1.ValidateMemberCredentialReq) (res *v1.ValidateMemberCredentialRes, err error) {
	return service.Auth().ValidateMemberCredential(ctx, req)
}
