package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	result, err := c.RoleRpcService.CreateRole(ctx, &oRole.CreateRoleReq{
		Code:   req.Code,
		Name:   req.Name,
		Status: req.Status,
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateRoleRes{
		Id: result.Id,
	}, nil
}
