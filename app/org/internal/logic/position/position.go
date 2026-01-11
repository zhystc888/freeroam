package position

import (
	"context"
	v1 "freeroam/app/org/api/position/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/do"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"
	"freeroam/common/tools/enum"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(&sPosition{})
}

// ListPosition 职务列表
func (s *sPosition) ListPosition(ctx context.Context, in *v1.ListPositionReq) (*v1.ListPositionRes, error) {
	m := dao.Position
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole

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
		// 获取关联的组织 ID
		orgIds, err := mPosOrg.Ctx(ctx).
			Where(mPosOrg.Columns().PositionId, pos.Id).
			Where(mPosOrg.Columns().IsDeleted, false).
			Fields(mPosOrg.Columns().OrgId).
			Array()
		if err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}

		// 获取关联的角色 ID
		roleIds, err := mPosRole.Ctx(ctx).
			Where(mPosRole.Columns().PositionId, pos.Id).
			Where(mPosRole.Columns().IsDeleted, false).
			Fields(mPosRole.Columns().RoleId).
			Array()
		if err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
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
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole

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

	// 获取关联的组织 ID
	orgIds, err := mPosOrg.Ctx(ctx).
		Where(mPosOrg.Columns().PositionId, pos.Id).
		Where(mPosOrg.Columns().IsDeleted, false).
		Fields(mPosOrg.Columns().OrgId).
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 获取关联的角色 ID
	roleIds, err := mPosRole.Ctx(ctx).
		Where(mPosRole.Columns().PositionId, pos.Id).
		Where(mPosRole.Columns().IsDeleted, false).
		Fields(mPosRole.Columns().RoleId).
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
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
			CreateBy:  0, // TODO: 从上下文获取用户ID
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
					mPosOrg.Columns().CreateBy:   0, // TODO: 从上下文获取用户ID
				})
			}
			_, err = tx.Model(mPosOrg.Table()).Ctx(ctx).Data(insertData).Insert()
			if err != nil {
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
					mPosRole.Columns().CreateBy:   0, // TODO: 从上下文获取用户ID
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

	return &v1.CreatePositionRes{
		Id: positionId,
	}, nil
}

// UpdatePosition 编辑职务
func (s *sPosition) UpdatePosition(ctx context.Context, in *v1.UpdatePositionReq) (*v1.UpdatePositionRes, error) {
	m := dao.Position
	mPosOrg := dao.PositionOrg
	mPosRole := dao.PositionRole

	// 检查职务是否存在
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
			UpdateBy: 0, // TODO: 从上下文获取用户ID
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
		// 先软删除旧的
		_, err = tx.Model(mPosOrg.Table()).Ctx(ctx).
			Where(mPosOrg.Columns().PositionId, in.Id).
			Data(g.Map{
				mPosOrg.Columns().IsDeleted: true,
				mPosOrg.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
				mPosOrg.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// 插入新的
		if len(in.OrgIds) > 0 {
			insertData := make([]g.Map, 0, len(in.OrgIds))
			for _, orgId := range in.OrgIds {
				insertData = append(insertData, g.Map{
					mPosOrg.Columns().PositionId: in.Id,
					mPosOrg.Columns().OrgId:      orgId,
					mPosOrg.Columns().CreateBy:   0, // TODO: 从上下文获取用户ID
				})
			}
			_, err = tx.Model(mPosOrg.Table()).Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return err
			}
		}

		// 覆盖写职务-角色关联
		// 先软删除旧的
		_, err = tx.Model(mPosRole.Table()).Ctx(ctx).
			Where(mPosRole.Columns().PositionId, in.Id).
			Data(g.Map{
				mPosRole.Columns().IsDeleted: true,
				mPosRole.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
				mPosRole.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// 插入新的
		if len(in.RoleIds) > 0 {
			insertData := make([]g.Map, 0, len(in.RoleIds))
			for _, roleId := range in.RoleIds {
				insertData = append(insertData, g.Map{
					mPosRole.Columns().PositionId: in.Id,
					mPosRole.Columns().RoleId:     roleId,
					mPosRole.Columns().CreateBy:   0, // TODO: 从上下文获取用户ID
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

	// 检查职务是否存在
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

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 软删除职务
		_, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().Id, in.Id).
			Data(g.Map{
				m.Columns().IsDeleted: true,
				m.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
				m.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// 软删除职务-组织关联
		_, err = tx.Model(mPosOrg.Table()).Ctx(ctx).
			Where(mPosOrg.Columns().PositionId, in.Id).
			Data(g.Map{
				mPosOrg.Columns().IsDeleted: true,
				mPosOrg.Columns().DeleteBy:  0,
				mPosOrg.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// 软删除职务-角色关联
		_, err = tx.Model(mPosRole.Table()).Ctx(ctx).
			Where(mPosRole.Columns().PositionId, in.Id).
			Data(g.Map{
				mPosRole.Columns().IsDeleted: true,
				mPosRole.Columns().DeleteBy:  0,
				mPosRole.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// TODO: 同步清理成员侧绑定（软删除）：member_org_position 中 position_id = id

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
		return nil, gerror.NewCode(berror.IncorrectParameters, "组织ID不能为空")
	}

	m := dao.Position
	mPosOrg := dao.PositionOrg

	// 查询关联了这些组织的职务ID
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
		code, err := enum.GetByTypeAndCode("position_status", "enable")
		if err != nil {
			return nil, gerror.WrapCode(berror.CodeInternal, err)
		}

		query = query.Where(m.Columns().Status, code.EnumValue)
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
