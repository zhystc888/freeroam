package position

import (
	"context"
	v1 "freeroam/app/org/api/position/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/do"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"
	"freeroam/common/consts/enum"
	"freeroam/common/tools/jwt_claims"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(&sPosition{})
}

// ListPosition 职务列表
func (s *sPosition) ListPosition(ctx context.Context, in *v1.ListPositionReq) (*v1.ListPositionRes, error) {
	m := dao.Position

	query := m.Ctx(ctx).Where(m.Columns().IsDeleted, false)

	if in.Keyword != "" {
		query = query.WhereLike(m.Columns().Name, "%"+in.Keyword+"%")
	}
	if in.Status != "" {
		query = query.Where(m.Columns().Status, in.Status)
	}

	// 获取总数
	total, err := query.Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 分页
	page := int(in.Page)
	pageSize := int(in.PageSize)

	var positions []*entity.Position
	err = query.OrderDesc(m.Columns().CreateAt).Page(page, pageSize).Scan(&positions)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	list := make([]*v1.PositionListItem, 0, len(positions))
	for _, pos := range positions {
		// 获取关联数据
		orgIds, roleIds, err := s.retrieveRelatedData(ctx, pos.Id)
		if err != nil {
			return nil, err
		}

		list = append(list, &v1.PositionListItem{
			Id:        int64(pos.Id),
			Name:      pos.Name,
			Status:    pos.Status,
			DataScope: pos.DataScope,
			OrgIds:    orgIds.Int64s(),
			RoleIds:   roleIds.Int64s(),
			CreateAt:  pos.CreateAt.Unix(),
		})
	}

	return &v1.ListPositionRes{
		List:     list,
		Total:    int64(total),
		Page:     int64(page),
		PageSize: int64(pageSize),
	}, nil
}

// GetPosition 获取职务详情
func (s *sPosition) GetPosition(ctx context.Context, in *v1.GetPositionReq) (*v1.GetPositionRes, error) {
	m := dao.Position

	var pos entity.Position
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&pos)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if pos.Id == 0 {
		return nil, gerror.NewCode(berror.PositionNotExist)
	}

	// 获取关联数据
	orgIds, roleIds, err := s.retrieveRelatedData(ctx, pos.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetPositionRes{
		Id:        int64(pos.Id),
		Name:      pos.Name,
		Status:    pos.Status,
		DataScope: pos.DataScope,
		OrgIds:    orgIds.Int64s(),
		RoleIds:   roleIds.Int64s(),
		CreateAt:  pos.CreateAt.Unix(),
		UpdateAt:  pos.UpdateAt.Unix(),
	}, nil
}

// CreatePosition 新建职务
func (s *sPosition) CreatePosition(ctx context.Context, in *v1.CreatePositionReq) (*v1.CreatePositionRes, error) {
	memberId := jwt_claims.GetMemberId(ctx)
	if memberId == 0 {
		return nil, berror.NewCode(berror.NotFindMemberIdFromCtx)
	}
	m := dao.Position
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole

	// 同一组织下职务名称重复校验
	if len(in.OrgIds) > 0 {
		err := s.checkPositionNameDuplicate(ctx, 0, in.Name, in.OrgIds)
		if err != nil {
			return nil, err
		}
	}

	var positionId int64
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建职务
		data := do.Position{
			Name:      in.Name,
			Status:    in.Status,
			DataScope: in.DataScope,
			CreateBy:  memberId,
		}

		result, err := tx.Model(m.Table()).Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}

		positionId, err = result.LastInsertId()
		if err != nil {
			return err
		}

		// 创建职务-组织关联
		if len(in.OrgIds) > 0 {
			insertData := make([]g.Map, 0, len(in.OrgIds))
			for _, orgId := range in.OrgIds {
				insertData = append(insertData, g.Map{
					mPosOrg.Columns().PositionId: positionId,
					mPosOrg.Columns().OrgId:      orgId,
					mPosOrg.Columns().CreateBy:   memberId,
				})
			}
			if _, err = tx.Model(mPosOrg.Table()).Ctx(ctx).Data(insertData).Insert(); err != nil {
				return err
			}
		}

		// 创建职务-角色关联
		if len(in.RoleIds) > 0 {
			insertData := make([]g.Map, 0, len(in.RoleIds))
			for _, roleId := range in.RoleIds {
				insertData = append(insertData, g.Map{
					mPosRole.Columns().PositionId: positionId,
					mPosRole.Columns().RoleId:     roleId,
					mPosRole.Columns().CreateBy:   memberId,
				})
			}
			if _, err = tx.Model(mPosRole.Table()).Ctx(ctx).Data(insertData).Insert(); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.CreatePositionRes{
		Id: positionId,
	}, nil
}

// UpdatePosition 编辑职务
func (s *sPosition) UpdatePosition(ctx context.Context, in *v1.UpdatePositionReq) (*v1.UpdatePositionRes, error) {
	m := dao.Position
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole
	memberId := jwt_claims.GetMemberId(ctx)
	if memberId == 0 {
		return nil, berror.NewCode(berror.NotFindMemberIdFromCtx)
	}
	// 检查职务是否存在
	exists, err := m.CheckExists(ctx, in.Id)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if !exists {
		return nil, gerror.NewCode(berror.PositionNotExist)
	}

	// 同一组织下职务名称重复校验
	if len(in.OrgIds) > 0 && in.Name != "" {
		err := s.checkPositionNameDuplicate(ctx, in.Id, in.Name, in.OrgIds)
		if err != nil {
			return nil, err
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新职务
		data := do.Position{
			UpdateBy: memberId,
		}
		if in.Name != "" {
			data.Name = in.Name
		}
		if in.Status != "" {
			data.Status = in.Status
		}
		if in.DataScope != "" {
			data.DataScope = in.DataScope
		}

		_, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().Id, in.Id).
			Data(data).
			Update()
		if err != nil {
			return err
		}

		// 覆盖写职务-组织关联
		// 先删除旧的
		if _, err = tx.Model(mPosOrg.Table()).Ctx(ctx).Unscoped().
			Where(mPosOrg.Columns().PositionId, in.Id).
			Delete(); err != nil {
			return err
		}

		// 插入新的
		if len(in.OrgIds) > 0 {
			insertData := make([]g.Map, 0, len(in.OrgIds))
			for _, orgId := range in.OrgIds {
				insertData = append(insertData, g.Map{
					mPosOrg.Columns().PositionId: in.Id,
					mPosOrg.Columns().OrgId:      orgId,
					mPosOrg.Columns().CreateBy:   memberId,
				})
			}
			_, err = tx.Model(mPosOrg.Table()).Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return err
			}
		}

		// 覆盖写职务-角色关联
		// 先删除旧的
		if _, err = tx.Model(mPosRole.Table()).Ctx(ctx).Unscoped().
			Where(mPosRole.Columns().PositionId, in.Id).
			Delete(); err != nil {
			return err
		}

		// 插入新的
		if len(in.RoleIds) > 0 {
			insertData := make([]g.Map, 0, len(in.RoleIds))
			for _, roleId := range in.RoleIds {
				insertData = append(insertData, g.Map{
					mPosRole.Columns().PositionId: in.Id,
					mPosRole.Columns().RoleId:     roleId,
					mPosRole.Columns().CreateBy:   memberId,
				})
			}
			_, err = tx.Model(mPosRole.Table()).Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.UpdatePositionRes{
		Success: true,
	}, nil
}

// DeletePosition 删除职务
func (s *sPosition) DeletePosition(ctx context.Context, in *v1.DeletePositionReq) (*v1.DeletePositionRes, error) {
	m := dao.Position
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole
	memberId := jwt_claims.GetMemberId(ctx)
	if memberId == 0 {
		return nil, berror.NewCode(berror.NotFindMemberIdFromCtx)
	}
	// 检查职务是否存在
	exists, err := m.CheckExists(ctx, in.Id)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if !exists {
		return nil, gerror.NewCode(berror.PositionNotExist)
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 软删除职务
		if err = m.DeleteById(ctx, in.Id, memberId, tx); err != nil {
			return berror.WrapCode(berror.DBErr, err)
		}

		// 软删除职务-组织关联
		if err = mPosOrg.DeleteByPositionId(ctx, in.Id, memberId, tx); err != nil {
			return berror.WrapCode(berror.DBErr, err)
		}

		// 软删除职务-角色关联
		if err = mPosRole.DeleteByPositionId(ctx, in.Id, memberId, tx); err != nil {
			return berror.WrapCode(berror.DBErr, err)
		}

		// TODO: 成员关联删除

		return nil
	})

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.DeletePositionRes{
		Success: true,
	}, nil
}

// GetPositionOptions 按组织获取可选职务集合
func (s *sPosition) GetPositionOptions(ctx context.Context, in *v1.GetPositionOptionsReq) (*v1.GetPositionOptionsRes, error) {
	if len(in.OrgIds) == 0 {
		return nil, gerror.NewCode(berror.DataNotExist, "组织 ID 不能为空")
	}

	m := dao.Position
	mPosOrg := dao.PositionOrg

	// 查询关联了这些组织的职务 ID
	positionIds, err := mPosOrg.Ctx(ctx).
		Where(mPosOrg.Columns().IsDeleted, false).
		WhereIn(mPosOrg.Columns().OrgId, in.OrgIds).
		Fields(mPosOrg.Columns().PositionId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	if len(positionIds) == 0 {
		return &v1.GetPositionOptionsRes{
			List: make([]*v1.PositionOption, 0),
		}, nil
	}

	// 查询职务
	query := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false).
		WhereIn(m.Columns().Id, positionIds)

	// 默认仅返回启用状态
	if in.Status != "" {
		query = query.Where(m.Columns().Status, in.Status)
	} else {
		query = query.Where(m.Columns().Status, enum.PositionStatusEnabled)
	}

	if in.Keyword != "" {
		query = query.WhereLike(m.Columns().Name, "%"+in.Keyword+"%")
	}

	var positions []*entity.Position
	err = query.Scan(&positions)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	list := make([]*v1.PositionOption, 0, len(positions))
	for _, pos := range positions {
		list = append(list, &v1.PositionOption{
			Id:        int64(pos.Id),
			Name:      pos.Name,
			Status:    pos.Status,
			DataScope: pos.DataScope,
		})
	}

	return &v1.GetPositionOptionsRes{
		List: list,
	}, nil
}

// checkPositionNameDuplicate 同一组织下职务名称重复校验
func (s *sPosition) checkPositionNameDuplicate(ctx context.Context, excludeId int64, name string, orgIds []int64) error {
	m := dao.Position
	mPosOrg := dao.PositionOrg

	// 对每个组织，查询该组织下所有关联职务，检查是否存在同名职务
	for _, orgId := range orgIds {
		// 获取该组织关联的所有职务 ID
		positionIds, err := mPosOrg.Ctx(ctx).
			Where(mPosOrg.Columns().OrgId, orgId).
			Where(mPosOrg.Columns().IsDeleted, false).
			Fields(mPosOrg.Columns().PositionId).
			Array()
		if err != nil {
			return gerror.NewCode(berror.DBErr, err.Error())
		}

		if len(positionIds) == 0 {
			continue
		}

		// 检查是否存在同名职务（排除自己）
		query := m.Ctx(ctx).
			WhereIn(m.Columns().Id, positionIds).
			Where(m.Columns().Name, name).
			Where(m.Columns().IsDeleted, false)

		if excludeId > 0 {
			query = query.WhereNot(m.Columns().Id, excludeId)
		}

		count, err := query.Count()
		if err != nil {
			return gerror.NewCode(berror.DBErr, err.Error())
		}
		if count > 0 {
			return gerror.NewCode(berror.PositionNameDuplicate)
		}
	}

	return nil
}

// 获取关联数据 return：关联组织id，关联角色id
func (s *sPosition) retrieveRelatedData(ctx context.Context, positionId uint64) (gdb.Array, gdb.Array, error) {
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole

	// 获取关联的组织 ID
	orgIds, err := mPosOrg.Ctx(ctx).
		Where(mPosOrg.Columns().PositionId, positionId).
		Where(mPosOrg.Columns().IsDeleted, false).
		Fields(mPosOrg.Columns().OrgId).
		Array()
	if err != nil {
		return nil, nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 获取关联的角色 ID
	roleIds, err := mPosRole.Ctx(ctx).
		Where(mPosRole.Columns().PositionId, positionId).
		Where(mPosRole.Columns().IsDeleted, false).
		Fields(mPosRole.Columns().RoleId).
		Array()
	if err != nil {
		return nil, nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	return orgIds, roleIds, nil
}
