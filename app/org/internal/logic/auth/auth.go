package auth

import (
	"context"

	v1 "freeroam/app/org/api/auth/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(&sAuth{})
}

// ValidateMemberCredential 校验成员账号密码
func (s *sAuth) ValidateMemberCredential(ctx context.Context, req *v1.ValidateMemberCredentialReq) (res *v1.ValidateMemberCredentialRes, err error) {
	m := dao.Member
	var member entity.Member
	// 查询成员信息（账号、未删除）
	err = m.Ctx(ctx).Safe(false).
		Where(m.Columns().Username, req.Username).
		Where(m.Columns().IsDeleted, false).
		Scan(&member)

	if err != nil {
		return nil, berror.NewCode(berror.DBErr, err.Error())
	}

	if member.Id == 0 {
		return nil, berror.NewCode(berror.CodeAccountNotExist)
	}

	// 校验密码（bcrypt）
	passwordMatch := false
	err = bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(req.Password))
	if err == nil {
		passwordMatch = true
	}

	if !passwordMatch {
		return nil, gerror.NewCode(berror.CodePasswordError)
	}

	// 校验是否禁用 TODO 用枚举
	if member.Status == 2 {
		return nil, berror.NewCode(berror.CodeAccountDisabled)
	}

	// 校验是否已离职（resigned_at 不为空） TODO 用枚举
	if member.Status == 3 {
		return nil, berror.NewCode(berror.CodeAccountResigned)
	}

	res = &v1.ValidateMemberCredentialRes{
		MemberId:  member.Id,
		Username:  member.Username,
		Name:      member.Name,
		Status:    uint32(member.Status),
		IsDeleted: uint32(member.IsDeleted),
	}

	// 如果已离职，设置 resigned_at
	if member.ResignedAt != nil && !member.ResignedAt.IsZero() {
		res.ResignedAt = timestamppb.New(member.ResignedAt.Time)
	}

	return res, nil
}
