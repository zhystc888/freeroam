package org

import (
	"context"
	"fmt"
	"strings"

	v1 "freeroam/app/org/api/org/v1"
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

type sOrg struct{}

func init() {
	service.RegisterOrg(&sOrg{})
}

// GetOrgTree 获取组织树
func (s *sOrg) GetOrgTree(ctx context.Context, in *v1.GetOrgTreeReq) (*v1.GetOrgTreeRes, error) {
	m := dao.Org

	// 查询所有未删除的组织
	query := m.Ctx(ctx).Where(m.Columns().IsDeleted, false)

	// 如果有关键字搜索，先找到匹配的组织及其子树
	if in.Keyword != "" {
		query = query.WhereLike(m.Columns().Name, "%"+in.Keyword+"%")
	}

	orgs := make([]*entity.Org, 0)
	if err := query.
		OrderAsc(m.Columns().ParentId).
		OrderAsc(m.Columns().Sort).
		Scan(&orgs); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 构建树结构
	tree := buildOrgTree(orgs, 0)

	return &v1.GetOrgTreeRes{
		List: tree,
	}, nil
}

// buildOrgTree 递归构建组织树
func buildOrgTree(orgs []*entity.Org, parentId uint64) []*v1.OrgTreeNode {
	// 1. 按 ParentId 分组
	childrenMap := make(map[uint64][]*v1.OrgTreeNode)
	for _, org := range orgs {
		node := &v1.OrgTreeNode{
			Id:          int64(org.Id),
			Name:        org.Name,
			FullName:    org.FullName,
			Code:        org.Code,
			Category:    org.Category,
			Status:      org.Status,
			MemberCount: 0,
			Children:    make([]*v1.OrgTreeNode, 0),
		}
		pid := org.ParentId
		childrenMap[pid] = append(childrenMap[pid], node)
	}

	// 2. 递归组装 (利用Map查找，复杂度 O(N))
	var attachChildren func(nodes []*v1.OrgTreeNode)
	attachChildren = func(nodes []*v1.OrgTreeNode) {
		for _, node := range nodes {
			if children, ok := childrenMap[uint64(node.Id)]; ok {
				node.Children = children
				node.MemberCount = int64(len(children))
				attachChildren(children)
			}
		}
	}

	// 3. 获取根节点列表并开始组装
	roots := childrenMap[parentId]
	attachChildren(roots)

	return roots
}

// ListOrg 获取组织列表
func (s *sOrg) ListOrg(ctx context.Context, in *v1.ListOrgReq) (*v1.ListOrgRes, error) {
	m := dao.Org

	query := m.Ctx(ctx).Where(m.Columns().IsDeleted, false)

	if in.Keyword != "" {
		query = query.WhereLike(m.Columns().Name, "%"+in.Keyword+"%")
	}
	if in.Code != "" {
		query = query.Where(m.Columns().Code, in.Code)
	}
	if in.Category != "" {
		query = query.Where(m.Columns().Category, in.Category)
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

	var orgs []*entity.Org
	if err = query.
		OrderDesc(m.Columns().CreateAt).
		Page(page, pageSize).
		Scan(&orgs); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	list := make([]*v1.OrgListItem, 0, len(orgs))
	for _, org := range orgs {
		list = append(list, &v1.OrgListItem{
			Id:       int64(org.Id),
			Name:     org.Name,
			FullName: org.FullName,
			Code:     org.Code,
			Status:   org.Status,
			Category: org.Category,
			CreateAt: org.CreateAt.Unix(),
		})
	}

	return &v1.ListOrgRes{
		List:     list,
		Total:    int64(total),
		Page:     int64(page),
		PageSize: int64(pageSize),
	}, nil
}

// GetOrg 获取组织详情
func (s *sOrg) GetOrg(ctx context.Context, in *v1.GetOrgReq) (*v1.GetOrgRes, error) {
	m := dao.Org

	var org entity.Org
	if err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&org); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if org.Id == 0 {
		return nil, gerror.NewCode(berror.OrgNotExist)
	}

	return &v1.GetOrgRes{
		Id:       int64(org.Id),
		ParentId: int64(org.ParentId),
		Name:     org.Name,
		FullName: org.FullName,
		Code:     org.Code,
		Category: org.Category,
		Status:   org.Status,
		Sort:     int32(org.Sort),
		Path:     org.Path,
		CreateAt: org.CreateAt.Unix(),
		UpdateAt: org.UpdateAt.Unix(),
	}, nil
}

// CreateOrg 新建组织
func (s *sOrg) CreateOrg(ctx context.Context, in *v1.CreateOrgReq) (*v1.CreateOrgRes, error) {
	m := dao.Org

	// 检查编码是否已存在
	exist, err := m.Ctx(ctx).
		Where(m.Columns().Code, in.Code).
		Where(m.Columns().IsDeleted, false).
		Exist()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if exist {
		return nil, gerror.NewCode(berror.OrgCodeAlreadyExist)
	}

	// 获取父级路径
	var parentPath string
	if in.ParentId > 0 {
		var parent entity.Org
		if err = m.Ctx(ctx).
			Where(m.Columns().Id, in.ParentId).
			Where(m.Columns().IsDeleted, false).
			Scan(&parent); err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}
		if parent.Id == 0 {
			return nil, gerror.NewCode(berror.OrgNotExist)
		}
		parentPath = parent.Path
	} else {
		parentPath = "/"
	}

	// 生成全称
	var fullName string
	if in.ParentId > 0 {
		var parent entity.Org
		_ = m.Ctx(ctx).Where(m.Columns().Id, in.ParentId).Scan(&parent)
		if parent.FullName != "" {
			fullName = parent.FullName + "/" + in.Name
		} else {
			fullName = in.Name
		}
	} else {
		fullName = in.Name
	}

	// 创建组织
	data := do.Org{
		ParentId: in.ParentId,
		Name:     in.Name,
		FullName: fullName,
		Code:     in.Code,
		Category: in.Category,
		Status:   in.Status,
		Sort:     in.Sort,
		Path:     parentPath, // 临时路径，插入后更新
		CreateBy: 0,          // TODO: 从上下文获取用户ID
	}

	result, err := m.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 更新路径
	newPath := fmt.Sprintf("%s%d/", parentPath, id)
	if _, err = m.Ctx(ctx).
		Where(m.Columns().Id, id).
		Data(g.Map{m.Columns().Path: newPath}).
		Update(); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.CreateOrgRes{
		Id: id,
	}, nil
}

// UpdateOrg 编辑组织
func (s *sOrg) UpdateOrg(ctx context.Context, in *v1.UpdateOrgReq) (*v1.UpdateOrgRes, error) {
	m := dao.Org

	// 1. 获取当前组织信息 (检查是否存在)
	var org entity.Org
	if err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&org); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if org.Id == 0 {
		return nil, gerror.NewCode(berror.OrgNotExist)
	}

	// 2. 检查编码唯一性 (如果修改了编码)
	if in.Code != "" && in.Code != org.Code {
		count, err := m.Ctx(ctx).
			Where(m.Columns().Code, in.Code).
			Where(m.Columns().IsDeleted, false).
			WhereNot(m.Columns().Id, in.Id).
			Count()
		if err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}
		if count > 0 {
			return nil, gerror.NewCode(berror.OrgCodeAlreadyExist)
		}
	}

	// 3. 计算变更后的核心字段: ParentId, Path, FullName
	targetParentId := int64(org.ParentId)
	targetPath := org.Path
	targetFullName := org.FullName

	// 判断是否涉及结构/名称变更
	isMove := in.ParentId != 0 && in.ParentId != int64(org.ParentId)
	isRename := in.Name != "" && in.Name != org.Name

	// 如果涉及移动或改名，可能需要查询父节点信息
	var parentOrg entity.Org
	needParentInfo := isMove || (isRename && targetParentId > 0)

	if needParentInfo {
		lookupId := targetParentId
		if isMove {
			lookupId = in.ParentId
		}

		if lookupId > 0 {
			if err := m.Ctx(ctx).
				Where(m.Columns().Id, lookupId).
				Where(m.Columns().IsDeleted, false).
				Scan(&parentOrg); err != nil {
				return nil, gerror.NewCode(berror.DBErr, err.Error())
			}
			if parentOrg.Id == 0 {
				return nil, gerror.NewCode(berror.OrgNotExist)
			}
		}
	}

	// 处理移动带来的 Path 和 ParentId 变更
	if isMove {
		// 防环校验
		if strings.HasPrefix(parentOrg.Path, org.Path) {
			return nil, gerror.NewCode(berror.OrgMoveIllegal)
		}
		targetParentId = in.ParentId

		// 计算新 Path
		if targetParentId == 0 {
			targetPath = fmt.Sprintf("/%d/", org.Id)
		} else {
			targetPath = fmt.Sprintf("%s%d/", parentOrg.Path, org.Id)
		}
	}

	// 处理移动或改名带来的 FullName 变更
	if isMove || isRename {
		newName := org.Name
		if in.Name != "" {
			newName = in.Name
		}

		if targetParentId == 0 {
			targetFullName = newName
		} else {
			if parentOrg.FullName != "" {
				targetFullName = parentOrg.FullName + "/" + newName
			} else {
				targetFullName = newName
			}
		}
	}

	// 4. 构建更新数据
	updateData := g.Map{
		m.Columns().UpdateBy: 0, // TODO: 从上下文获取用户ID
	}
	if in.Name != "" {
		updateData[m.Columns().Name] = in.Name
	}
	if in.Code != "" {
		updateData[m.Columns().Code] = in.Code
	}
	if in.Category != "" {
		updateData[m.Columns().Category] = in.Category
	}
	if in.Status != "" {
		updateData[m.Columns().Status] = in.Status
	}
	if in.Sort >= 0 {
		updateData[m.Columns().Sort] = in.Sort
	}
	if isMove {
		updateData[m.Columns().ParentId] = targetParentId
		updateData[m.Columns().Path] = targetPath
	}
	if isMove || isRename {
		updateData[m.Columns().FullName] = targetFullName
	}

	// 5. 事务执行
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新自身
		_, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().Id, in.Id).
			Data(updateData).
			Update()
		if err != nil {
			return err
		}

		// 如果 Path 或 FullName 变更，递归更新子孙节点
		if targetPath != org.Path || targetFullName != org.FullName {
			if err = updateRecursiveInfo(tx, org.Id, org.Path, targetPath, org.FullName, targetFullName); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.UpdateOrgRes{Success: true}, nil
}

// DeleteOrg 删除组织
func (s *sOrg) DeleteOrg(ctx context.Context, in *v1.DeleteOrgReq) (*v1.DeleteOrgRes, error) {
	m := dao.Org

	// 检查组织是否存在
	var org entity.Org
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, false).
		Scan(&org)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if org.Id == 0 {
		return nil, gerror.NewCode(berror.OrgNotExist)
	}

	// 使用事务删除组织及其子树
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 获取所有子组织
		var orgs []*entity.Org
		if err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().IsDeleted, false).
			WhereLike(m.Columns().Path, org.Path+"%").
			Scan(&orgs); err != nil {
			return err
		}

		// 需要删除的组织和关联关系
		positionOrg := dao.PositionOrg
		for _, item := range orgs {
			if _, err = tx.Model(m.Table()).Ctx(ctx).
				Where(m.Columns().Id, item.Id).
				Data(g.Map{
					m.Columns().IsDeleted: true,
					m.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
					m.Columns().DeletedAt: gtime.Now(),
				}).
				Update(); err != nil {
				return err
			}

			// 清理关联
			if _, err = tx.Model(positionOrg.Table()).Ctx(ctx).
				Where(positionOrg.Columns().IsDeleted, false).
				Where(positionOrg.Columns().OrgId, item.Id).
				Data(g.Map{
					m.Columns().IsDeleted: true,
					m.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
					m.Columns().DeletedAt: gtime.Now(),
				}).
				Update(); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.DeleteOrgRes{
		Success: true,
	}, nil
}

// MoveOrg 拖拽移动/排序
func (s *sOrg) MoveOrg(ctx context.Context, in *v1.MoveOrgReq) (*v1.MoveOrgRes, error) {
	m := dao.Org

	// 检查组织是否存在
	var org entity.Org
	if err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, 0).
		Scan(&org); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if org.Id == 0 {
		return nil, gerror.NewCode(berror.OrgNotExist)
	}

	// 只更新排序
	if in.NewParentId == int64(org.ParentId) {
		if _, err := m.Ctx(ctx).
			Where(m.Columns().Id, org.Id).
			Data(g.Map{m.Columns().Sort: in.NewSort}).
			Update(); err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}

		return &v1.MoveOrgRes{
			Success: true,
		}, nil
	}

	// 获取新父级路径和全称
	var newParentPath string
	var newParentFullName string
	if in.NewParentId > 0 {
		var newParent entity.Org
		if err := m.Ctx(ctx).
			Where(m.Columns().Id, in.NewParentId).
			Where(m.Columns().IsDeleted, 0).
			Scan(&newParent); err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}
		if newParent.Id == 0 {
			return nil, gerror.NewCode(berror.OrgNotExist)
		}
		// 防环校验：不允许移动到自己的子树
		if strings.HasPrefix(newParent.Path, org.Path) {
			return nil, gerror.NewCode(berror.OrgMoveIllegal)
		}
		newParentPath = newParent.Path
		newParentFullName = newParent.FullName
	} else {
		newParentPath = "/"
		newParentFullName = ""
	}

	// 计算新路径
	newPath := fmt.Sprintf("%s%d/", newParentPath, org.Id)
	oldPath := org.Path

	// 计算新全称
	newFullName := org.Name
	if newParentFullName != "" {
		newFullName = newParentFullName + "/" + org.Name
	}
	oldFullName := org.FullName

	// 使用事务更新
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新当前节点
		if _, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().Id, in.Id).
			Data(g.Map{
				m.Columns().ParentId: in.NewParentId,
				m.Columns().Sort:     in.NewSort,
				m.Columns().Path:     newPath,
				m.Columns().FullName: newFullName,
				m.Columns().UpdateBy: 0, // TODO: 从上下文获取用户ID
			}).
			Update(); err != nil {
			return err
		}

		// 批量更新子孙节点的路径和全称
		if err := updateRecursiveInfo(tx, org.Id, oldPath, newPath, oldFullName, newFullName); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.MoveOrgRes{
		Success: true,
	}, nil
}

// updateRecursiveInfo 递归更新子节点信息（Path和FullName）
func updateRecursiveInfo(tx gdb.TX, orgId uint64, oldPath, newPath, oldFullName, newFullName string) error {
	m := dao.Org
	// 如果路径有变化，批量更新子孙节点路径
	if oldPath != newPath {
		_, err := tx.Model(m.Table()).Data(g.Map{
			m.Columns().Path: gdb.Raw(fmt.Sprintf("REPLACE(%s, %s, %s)", oldPath, oldFullName, newFullName)),
		}).WhereLike(m.Columns().Path, oldPath+"%").
			Where(m.Columns().IsDeleted, "=", false).
			Where(m.Columns().Id, "<>", orgId).
			Update()
		if err != nil {
			return err
		}
	}

	// 如果全称有变化，批量更新子孙节点全称
	if oldFullName != newFullName {
		// 注意这里匹配 oldFullName + "/" 确保只替换作为前缀的部分
		if _, err := tx.Exec(
			fmt.Sprintf("UPDATE %s SET %s = REPLACE(%s, ?, ?) WHERE %s LIKE ? AND %s = 0 AND %s != ?",
				m.Table(), m.Columns().FullName, m.Columns().FullName, m.Columns().FullName, m.Columns().IsDeleted, m.Columns().Id),
			oldFullName, newFullName, oldFullName+"/%", orgId); err != nil {
			return err
		}
	}
	return nil
}
