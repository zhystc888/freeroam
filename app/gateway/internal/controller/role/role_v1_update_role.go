package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	result, err := c.RoleRpcService.UpdateRole(ctx, &oRole.UpdateRoleReq{
		Id:     req.Id,
		Name:   req.Name,
		Status: req.Status,
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRoleRes{
		Success: result.Success,
	}, nil
}
