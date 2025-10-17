package org

import (
	"context"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"
)

type sOrg struct{}

func init() {
	service.RegisterOrg(&sOrg{})
}

func (s *sOrg) Get(ctx context.Context, id int64) (res *model.OrgVo, err error) {
	m := dao.OrgStructure
	err = m.Ctx(ctx).WithAll().Where(m.Columns().Id, id).Scan(&res)
	// 封装error为制定异常信息
	if err != nil {
		err = berror.NewInternalError(err)
	}
	return
}

func (s *sOrg) GetList(ctx context.Context, params *model.OrgListDto) (res *model.ListReq[model.OrgListVo], err error) {
	m := dao.OrgStructure
	query := m.Ctx(ctx).Fields(res).Safe(false).Limit(params.GetOffset(), params.GetLimit())
	if params.Name != "" {
		query.WhereLike(m.Columns().Name, params.Name)
	}
	if params.ParentId != nil {
		query.Where(m.Columns().ParentId, params.ParentId)
	}
	if params.Code != "" {
		query.Where(m.Columns().Code, params.Code)
	}
	if params.Type != 0 {
		query.Where(m.Columns().Type, params.Type)
	}
	list := make([]*model.OrgListVo, 0)
	Total := 0
	res = &model.ListReq[model.OrgListVo]{
		List:  list,
		Total: Total,
	}
	err = query.ScanAndCount(&res.List, &res.Total, false)
	if err != nil {
		err = berror.NewInternalError(err)
	}

	return
}
