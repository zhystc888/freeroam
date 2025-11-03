package org

import (
	"context"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *sOrg) Create(ctx context.Context, dto *model.CreateOrgDto) error {
	m := dao.OrgStructure
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return berror.NewInternalError(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err := m.Ctx(ctx).TX(tx).Data(dto).InsertAndGetId()
	if err != nil {
		return berror.NewInternalError(err)
	}

	// 使用批量插入方法
	err = dao.OrgSupervisor.BatchInsertSupervisors(ctx, tx, id, dto.SupervisorIds)
	if err != nil {
		return berror.NewInternalError(err)
	}
	return nil
}

func (s *sOrg) Update(ctx context.Context, dto *model.UpdateOrgDto) error {
	m := dao.OrgStructure
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return berror.NewInternalError(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = m.Ctx(ctx).TX(tx).Where(m.Columns().Id, dto.Id).Data(dto).Update()
	if err != nil {
		return berror.NewInternalError(err)
	}

	// 使用批量插入方法
	err = dao.OrgSupervisor.BatchInsertSupervisors(ctx, tx, dto.Id, dto.SupervisorIds)
	if err != nil {
		return berror.NewInternalError(err)
	}
	return nil
}
