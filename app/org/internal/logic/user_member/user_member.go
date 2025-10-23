package user_member

import (
	"context"
	v1 "freeroam/app/org/api/user_member/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

type sUserMember struct{}

func init() {
	service.RegisterUserMember(&sUserMember{})
}

func (s *sUserMember) GetList(ctx context.Context, params *model.UserMemberListDto) (res *model.ListReq[model.UserMemberListVo], err error) {
	m := dao.UserMember
	query := m.Ctx(ctx).Fields(res).Safe(false).Limit(params.GetOffset(), params.GetLimit())
	if params.Username != "" {
		query.WhereLike(m.Columns().Username, params.Username+"%")
	}
	if params.Name != "" {
		query.WhereLike(m.Columns().Name, params.Name+"%")
	}
	if params.Mobile != "" {
		query.Where(m.Columns().Mobile, params.Mobile)
	}
	if params.Gender != nil {
		query.Where(m.Columns().Gender, params.Gender)
	}
	if params.Status != nil {
		query.Where(m.Columns().Status, params.Status)
	}
	list := make([]*model.UserMemberListVo, 0)
	Total := 0
	res = &model.ListReq[model.UserMemberListVo]{
		List:  list,
		Total: Total,
	}
	err = query.ScanAndCount(&res.List, &res.Total, false)
	if err != nil {
		err = berror.NewInternalError(err)
	}
	return
}

func (s *sUserMember) GetOne(ctx context.Context, userId int64) (res *model.UserMemberVo, err error) {
	m := dao.UserMember
	err = m.Ctx(ctx).WithAll().Where(m.Columns().UserId, userId).Scan(&res)
	// 封装error为制定异常信息
	if err != nil {
		err = berror.NewInternalError(err)
	}
	return
}

func (s *sUserMember) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	var userBase entity.UserBase
	m := dao.UserBase
	err = m.Ctx(ctx).Where(m.Columns().Id, req.UserId).Scan(&userBase)

	if userBase.Id == 0 {
		return res, berror.NewBizError(berror.DataNotExist, "用户不存在！")
	}

	// 验证原密码
	if err := bcrypt.CompareHashAndPassword([]byte(userBase.Password), []byte(req.OriginalPassword)); err != nil {
		return res, berror.NewBizError(berror.PasswordErr, "原密码不正确！")
	}

	// 生成新密码
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return res, berror.NewInternalError(err)
	}

	// 更新数据
	_, err = m.Ctx(ctx).Where(m.Columns().Id, req.UserId).Data(m.Columns().Password, gconv.String(pwd)).Update()
	if err != nil {
		return res, berror.NewInternalError(err)
	}

	return
}
