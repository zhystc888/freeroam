package admin_member

import (
	"context"

	"bbk/app/gateway/api/admin_member/v1"
	uam "bbk/app/user/api/admin_member/v1"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	result, err := c.UserRpcService.Get(ctx, &uam.GetReq{
		UserId: req.UserID,
	})

	if result != nil && result.UserId > 0 {
		res = &v1.GetRes{
			Username: result.Username,
		}
		res.UserID = result.UserId
		res.Name = result.Name
		res.Status = result.Status
	}

	return
}
