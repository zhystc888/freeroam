package permission

import (
	"context"

	"freeroam/app/gateway/api/permission/v1"
	sPermission "freeroam/app/system/api/permission/v1"
	"freeroam/common/tools/jwt_claims"
)

func (c *ControllerV1) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error) {
	// 从 JWT 中获取成员ID
	memberId := jwt_claims.GetMemberId(ctx)

	rpcReq := &sPermission.GetFrontPermissionsReq{
		MemberId: int64(memberId),
	}

	rpcRes, err := c.PermissionRpcService.GetFrontPermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GetFrontPermissionsRes{
		Permissions: rpcRes.Permissions,
	}, nil
}
