package user_member

import (
	"context"
	sUserMember "freeroam/app/org/api/user_member/v1"

	"freeroam/app/gateway/api/user_member/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	result, err := c.UserMemberRpcService.GetOne(ctx, &sUserMember.GetOneReq{UserId: req.UserId})
	if result != nil && result.UserId > 0 {
		res = &v1.GetOneRes{
			UserId:   result.UserId,
			Username: result.Username,
			Name:     result.Name,
			Mobile:   result.Mobile,
			Gender:   result.Gender,
			Status:   result.Status,
			Super:    result.Super,
		}
	}
	return
}
