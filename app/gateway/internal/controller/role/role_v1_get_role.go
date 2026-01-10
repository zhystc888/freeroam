package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) GetRole(ctx context.Context, req *v1.GetRoleReq) (res *v1.GetRoleRes, err error) {
	result, err := c.RoleRpcService.GetRole(ctx, &oRole.GetRoleReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetRoleRes{
		Id:       result.Id,
		Code:     result.Code,
		Name:     result.Name,
		Status:   result.Status,
		IsSystem: result.IsSystem,
		Remark:   result.Remark,
		CreateAt: result.CreateAt,
		UpdateAt: result.UpdateAt,
	}, nil
}
