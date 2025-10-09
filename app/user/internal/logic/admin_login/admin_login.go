package admin_login

import (
	v1 "bbk/app/user/api/admin_login/v1"
	"bbk/app/user/internal/dao"
	"bbk/app/user/internal/model"
	"bbk/app/user/internal/service"
	"bbk/common/berror"
	"bbk/common/enum"
	"context"
)

type sAdminLogin struct{}

func init() {
	service.RegisterAdminLogin(&sAdminLogin{})
}

func (s *sAdminLogin) AdminLogin(ctx context.Context, dto *model.LoginDto) (res *v1.LoginRes, err error) {
	da := dao.AdminUser
	adminInfo := &model.LoginAdminUserInfo{}

	db := da.Ctx(ctx).Safe(false).Fields(adminInfo)
	err = db.Where(da.Columns().Username, dto.Username).Scan(&adminInfo)
	if err != nil {
		return nil, berror.NewInternalError(err)
	}

	if adminInfo.UserID == 0 {
		return nil, berror.NewBizError(berror.CodeLoginVerificationFailed, "用户不存在")
	}

	if adminInfo.User.ID == 0 {
		return nil, berror.NewBizError(berror.CodeLoginVerificationFailed, "用户不存在")
	}

	if adminInfo.Status != int64(enum.Enabled) {
		return nil, berror.NewBizError(berror.CodeUserNotAvailable, "用户已禁用或未启用")
	}

	return
}
