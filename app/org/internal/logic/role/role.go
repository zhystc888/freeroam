package role

import (
	"context"
	v1 "freeroam/app/org/api/role/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/do"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sRole struct{}

func init() {
	service.RegisterRole(&sRole{})
}

// CreateRole 创建角色
func (s *sRole) CreateRole(ctx context.Context, in *v1.CreateRoleReq) (*v1.CreateRoleRes, error) {
	m := dao.Role

	// 检查角色编码是否已存在
	count, err := m.Ctx(ctx).
		Where(m.Columns().Code, in.Code).
		Where(m.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count > 0 {
		return nil, gerror.NewCode(berror.RoleCodeAlreadyExists)
	}

	// 创建角色
	data := do.Role{
		Code:     in.Code,
		Name:     in.Name,
		Status:   in.Status,
		Remark:   in.Remark,
		CreateBy: 0, // TODO: 从上下文获取用户ID
	}

	result, err := m.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.CreateRoleRes{
		Id: id,
	}, nil
}

// UpdateRole 更新角色
func (s *sRole) UpdateRole(ctx context.Context, in *v1.UpdateRoleReq) (*v1.UpdateRoleRes, error) {
	m := dao.Role

	// 检查角色是否存在
	count, err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	// 更新角色
	data := do.Role{
		UpdateBy: 0, // TODO: 从上下文获取用户ID
	}
	if in.Name != "" {
		data.Name = in.Name
	}
	if in.Status != "" {
		data.Status = in.Status
	}
	if in.Remark != "" {
		data.Remark = in.Remark
	}

	_, err = m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Data(data).
		Update()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.UpdateRoleRes{
		Success: true,
	}, nil
}

// DeleteRole 删除角色
func (s *sRole) DeleteRole(ctx context.Context, in *v1.DeleteRoleReq) (*v1.DeleteRoleRes, error) {
	m := dao.Role

	// 检查角色是否存在
	var role entity.Role
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&role)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if role.Id == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	// 检查是否为系统内置角色
	if role.IsSystem == 1 {
		return nil, gerror.NewCode(berror.BuiltInSystemRolesCannotDeleted)
	}

	// 检查角色是否被绑定
	positionRole := dao.PositionRole
	count, err := positionRole.Ctx(ctx).
		Where(positionRole.Columns().RoleId, in.Id).
		Where(positionRole.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count > 0 {
		return nil, gerror.NewCode(berror.RoleIsBoundCannotDeleted)
	}

	// 软删除
	_, err = m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Data(do.Role{
			IsDeleted: true,
			DeleteBy:  0, // TODO: 从上下文获取用户ID
			DeletedAt: gtime.Now(),
		}).
		Update()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.DeleteRoleRes{
		Success: true,
	}, nil
}

// GetRole 获取角色详情
func (s *sRole) GetRole(ctx context.Context, in *v1.GetRoleReq) (*v1.GetRoleRes, error) {
	m := dao.Role

	var role entity.Role
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&role)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if role.Id == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	return &v1.GetRoleRes{
		Id:       int64(role.Id),
		Code:     role.Code,
		Name:     role.Name,
		Status:   role.Status,
		IsSystem: role.IsSystem == 1,
		Remark:   role.Remark,
		CreateAt: role.CreateAt.Unix(),
		UpdateAt: role.UpdateAt.Unix(),
	}, nil
}

// ListRole 获取角色列表
func (s *sRole) ListRole(ctx context.Context, in *v1.ListRoleReq) (*v1.ListRoleRes, error) {
	m := dao.Role

	// 构建查询条件
	query := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false)

	if in.Keyword != "" {
		query = query.WhereLike(m.Columns().Code, "%"+in.Keyword+"%").
			WhereOrLike(m.Columns().Name, "%"+in.Keyword+"%")
	}
	if in.Status != "" {
		query = query.Where(m.Columns().Status, in.Status)
	}
	if in.IsSystem > 0 {
		query = query.Where(m.Columns().IsSystem, in.IsSystem == 1)
	}

	// 获取总数
	total, err := query.Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 分页查询
	var roles []*entity.Role
	page := int(in.Page)
	pageSize := int(in.PageSize)

	err = query.
		OrderDesc(m.Columns().Id).
		Page(page, pageSize).
		Scan(&roles)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 转换为响应格式
	list := make([]*v1.GetRoleRes, 0, len(roles))
	for _, role := range roles {
		list = append(list, &v1.GetRoleRes{
			Id:       int64(role.Id),
			Code:     role.Code,
			Name:     role.Name,
			Status:   role.Status,
			IsSystem: role.IsSystem == 1,
			Remark:   role.Remark,
			CreateAt: role.CreateAt.Unix(),
			UpdateAt: role.UpdateAt.Unix(),
		})
	}

	return &v1.ListRoleRes{
		List:     list,
		Total:    int64(total),
		Page:     int64(page),
		PageSize: int64(pageSize),
	}, nil
}

// GetRolePositionList 查询角色绑定的职务列表
func (s *sRole) GetRolePositionList(ctx context.Context, in *v1.GetRolePositionListReq) (*v1.GetRolePositionListRes, error) {
	// 检查角色是否存在
	m := dao.Role
	var role entity.Role
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.RoleId).
		Where(m.Columns().IsDeleted, false).
		Scan(&role)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if role.Id == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	// 查询角色绑定的职务列表
	positionRole := dao.PositionRole
	positionIds, err := positionRole.Ctx(ctx).
		Where(positionRole.Columns().RoleId, in.RoleId).
		Where(positionRole.Columns().IsDeleted, false).
		Fields(positionRole.Columns().PositionId).
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 查询获取职务信息
	page := int(in.Page)
	pageSize := int(in.PageSize)

	positions := make([]*entity.Position, 0, pageSize)
	position := dao.Position

	positionQuery := position.Ctx(ctx).
		WhereIn(position.Columns().Id, positionIds).
		Where(position.Columns().IsDeleted, false)

	if in.Keyword != "" {
		positionQuery = positionQuery.WhereLike(position.Columns().Name, "%"+in.Keyword+"%")
	}

	count, err := positionQuery.Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	if err = positionQuery.OrderAsc(position.Columns().Id).
		Page(page, pageSize).
		Scan(&positions); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 转换为响应格式
	list := make([]*v1.PositionItem, 0, pageSize)
	for _, item := range positions {
		list = append(list, &v1.PositionItem{
			PositionId:     int64(item.Id),
			PositionName:   item.Name,
			PositionStatus: item.Status,
			CreateAt:       item.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetRolePositionListRes{
		List:     list,
		Total:    int64(count),
		Page:     int64(page),
		PageSize: int64(pageSize),
	}, nil
}

// BatchAssignRolePosition 批量绑定职务到角色（覆盖式）
func (s *sRole) BatchAssignRolePosition(ctx context.Context, in *v1.BatchAssignRolePositionReq) (*v1.BatchAssignRolePositionRes, error) {
	// 检查角色是否存在
	m := dao.Role
	count, err := m.Ctx(ctx).
		Where(m.Columns().Id, in.RoleId).
		Where(m.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	// 使用事务处理
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 先删除该角色的所有职务关联
		positionRole := dao.PositionRole
		_, err := tx.Model(positionRole.Table()).Ctx(ctx).Unscoped().
			Where(positionRole.Columns().RoleId, in.RoleId).
			Data(g.Map{
				positionRole.Columns().IsDeleted: true,
				positionRole.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
				positionRole.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// 如果有新的职务ID列表，则插入新的关联
		if len(in.PositionIds) > 0 {
			// 验证职务是否存在
			position := dao.Position
			positionCount, err := tx.Model(position.Table()).Ctx(ctx).
				WhereIn(position.Columns().Id, in.PositionIds).
				Where(position.Columns().IsDeleted, false).
				Count()
			if err != nil {
				return err
			}
			if positionCount != len(in.PositionIds) {
				return gerror.NewCode(berror.DataNotExist, "部分职务不存在")
			}

			// 批量插入新的关联关系
			insertData := make([]g.Map, 0, len(in.PositionIds))
			for _, positionId := range in.PositionIds {
				insertData = append(insertData, g.Map{
					positionRole.Columns().RoleId:     in.RoleId,
					positionRole.Columns().PositionId: positionId,
					positionRole.Columns().CreateBy:   0, // TODO: 从上下文获取用户ID
				})
			}

			_, err = tx.Model(positionRole.Table()).Ctx(ctx).
				Data(insertData).
				Insert()
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.BatchAssignRolePositionRes{
		Success: true,
	}, nil
}

// GetRolePositionIds 查询角色当前绑定职务ID集合（表单回显）
func (s *sRole) GetRolePositionIds(ctx context.Context, in *v1.GetRolePositionIdsReq) (*v1.GetRolePositionIdsRes, error) {
	// 检查角色是否存在
	m := dao.Role
	count, err := m.Ctx(ctx).
		Where(m.Columns().Id, in.RoleId).
		Where(m.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count == 0 {
		return nil, gerror.NewCode(berror.RoleNotExist)
	}

	// 查询角色绑定的职务ID列表
	positionRole := dao.PositionRole
	positionIds, err := positionRole.Ctx(ctx).
		Where(positionRole.Columns().RoleId, in.RoleId).
		Where(positionRole.Columns().IsDeleted, false).
		Fields(positionRole.Columns().PositionId).
		OrderAsc(positionRole.Columns().PositionId).
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.GetRolePositionIdsRes{
		PositionIds: positionIds.Int64s(),
	}, nil
}
