package permission

import (
	"context"
	"freeroam/common/tools/enum"

	v1 "freeroam/app/gateway/api/permission/v1"
	oPermission "freeroam/app/org/api/permission/v1"
	"freeroam/common/tools/jwt_claims"
)

func (c *ControllerV1) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error) {
	// 从 JWT 中获取成员ID
	memberId := jwt_claims.GetMemberId(ctx)

	typeList, err := enum.GetByType("permissions_type")
	if err != nil {
		return nil, err
	}

	permTypes := make([]string, 0, len(typeList.Options)-1)
	for _, item := range typeList.Options {
		if item.EnumCode == "interface" {
			continue
		}
		permTypes = append(permTypes, item.EnumValue)
	}

	rpcReq := &oPermission.GetMemberPermissionsReq{
		MemberId:  int64(memberId),
		PermTypes: permTypes,
	}

	rpcRes, err := c.PermissionRpcService.GetMemberPermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GetFrontPermissionsRes{
		Permissions: rpcRes.Permissions,
	}, nil
}
