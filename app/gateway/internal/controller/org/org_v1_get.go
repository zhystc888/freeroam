package org

import (
	"bbk/app/gateway/api/org/v1"
	"context"

	sorg "bbk/app/org/api/org/v1"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	result, err := c.OrgRpcService.Get(ctx, &sorg.GetReq{Id: req.Id})
	if result != nil && result.Id > 0 {
		res = &v1.GetRes{
			Id:       result.Id,
			ParentId: result.ParentId,
			Name:     result.Name,
			Code:     result.Code,
			Type:     result.Type,
			Status:   result.Status,
		}

		res.Supervisors = make([]*v1.Supervisor, len(result.Supervisors))
		for k, v := range result.Supervisors {
			res.Supervisors[k] = &v1.Supervisor{
				UserId: v.GetId(),
				Name:   v.GetName(),
			}
		}
	}
	return
}
