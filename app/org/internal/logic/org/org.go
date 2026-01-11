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

	var orgs []*entity.Org
	err := query.OrderAsc(m.Columns().ParentId).OrderAsc(m.Columns().Sort).Scan(&orgs)
	if err != nil {
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
	var nodes []*v1.OrgTreeNode
	for _, org := range orgs {
		if org.ParentId == parentId {
			node := &v1.OrgTreeNode{
				Id:          int64(org.Id),
				Name:        org.Name,
				FullName:    org.FullName,
				Code:        org.Code,
				Category:    org.Category,
				Status:      org.Status,
				MemberCount: 0, // TODO: 后续实现成员统计
				Children:    buildOrgTree(orgs, org.Id),
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
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
	err = query.OrderDesc(m.Columns().CreateAt).Page(page, pageSize).Scan(&orgs)
	if err != nil {
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
	count, err := m.Ctx(ctx).
		Where(m.Columns().Code, in.Code).
		Where(m.Columns().IsDeleted, false).
		Count()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if count > 0 {
		return nil, gerror.NewCode(berror.OrgCodeAlreadyExist)
	}

	// 获取父级路径
	var parentPath string
	if in.ParentId > 0 {
		var parent entity.Org
		err := m.Ctx(ctx).
			Where(m.Columns().Id, in.ParentId).
			Where(m.Columns().IsDeleted, false).
			Scan(&parent)
		if err != nil {
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
	fullName := in.FullName
	if fullName == "" {
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
	_, err = m.Ctx(ctx).
		Where(m.Columns().Id, id).
		Data(g.Map{m.Columns().Path: newPath}).
		Update()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.CreateOrgRes{
		Id: id,
	}, nil
}

// UpdateOrg 编辑组织
func (s *sOrg) UpdateOrg(ctx context.Context, in *v1.UpdateOrgReq) (*v1.UpdateOrgRes, error) {
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

	// 检查编码是否重复（排除自己）
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
	// TODO: parentId 不是空 更新path等路径

	// 更新数据
	data := do.Org{
		UpdateBy: 0, // TODO: 从上下文获取用户ID
	}
	if in.Name != "" {
		data.Name = in.Name
	}
	if in.FullName != "" {
		data.FullName = in.FullName
	}
	if in.Code != "" {
		data.Code = in.Code
	}
	if in.Category != "" {
		data.Category = in.Category
	}
	if in.Status != "" {
		data.Status = in.Status
	}
	if in.Sort >= 0 {
		data.Sort = in.Sort
	}

	_, err = m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Data(data).
		Update()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.UpdateOrgRes{
		Success: true,
	}, nil
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
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 软删除该组织及其所有子组织
		_, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().IsDeleted, false).
			WhereLike(m.Columns().Path, org.Path+"%").
			Data(g.Map{
				m.Columns().IsDeleted: true,
				m.Columns().DeleteBy:  0, // TODO: 从上下文获取用户ID
				m.Columns().DeletedAt: gtime.Now(),
			}).
			Update()
		if err != nil {
			return err
		}

		// TODO: 清理成员关联（软删除）
		// member_org_unit 和 member_org_position 中涉及该组织及其子组织的记录

		return nil
	})

	if err != nil {
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
	err := m.Ctx(ctx).
		Where(m.Columns().Id, in.Id).
		Where(m.Columns().IsDeleted, 0).
		Scan(&org)
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}
	if org.Id == 0 {
		return nil, gerror.NewCode(berror.OrgNotExist)
	}

	// 获取新父级路径
	var newParentPath string
	if in.NewParentId > 0 {
		var newParent entity.Org
		err := m.Ctx(ctx).
			Where(m.Columns().Id, in.NewParentId).
			Where(m.Columns().IsDeleted, 0).
			Scan(&newParent)
		if err != nil {
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
	} else {
		newParentPath = "/"
	}

	// 计算新路径
	newPath := fmt.Sprintf("%s%d/", newParentPath, org.Id)
	oldPath := org.Path

	// 使用事务更新
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新当前节点
		_, err := tx.Model(m.Table()).Ctx(ctx).
			Where(m.Columns().Id, in.Id).
			Data(g.Map{
				m.Columns().ParentId: in.NewParentId,
				m.Columns().Sort:     in.NewSort,
				m.Columns().Path:     newPath,
				m.Columns().UpdateBy: 0, // TODO: 从上下文获取用户ID
			}).
			Update()
		if err != nil {
			return err
		}

		// 批量更新子孙节点的路径（旧前缀替换新前缀）
		if oldPath != newPath {
			_, err = tx.Exec(
				fmt.Sprintf("UPDATE %s SET path = REPLACE(path, ?, ?) WHERE path LIKE ? AND is_deleted = 0 AND id != ?",
					m.Table()),
				oldPath, newPath, oldPath+"%", in.Id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.MoveOrgRes{
		Success: true,
	}, nil
}
