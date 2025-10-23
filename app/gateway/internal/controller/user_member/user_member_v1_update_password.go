package user_member

import (
	"context"

	"freeroam/app/gateway/api/user_member/v1"
	sUserMember "freeroam/app/org/api/user_member/v1"
)

func (c *ControllerV1) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	_, err = c.UserMemberRpcService.UpdatePassword(ctx, &sUserMember.UpdatePasswordReq{
		UserId:           req.UserId,
		OriginalPassword: req.OriginalPassword,
		NewPassword:      req.NewPassword,
	})

	return
}
