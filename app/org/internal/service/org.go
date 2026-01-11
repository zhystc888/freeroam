// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/org/api/org/v1"
)

type (
	IOrg interface {
		// GetOrgTree 获取组织树
		GetOrgTree(ctx context.Context, in *v1.GetOrgTreeReq) (*v1.GetOrgTreeRes, error)
		// ListOrg 获取组织列表
		ListOrg(ctx context.Context, in *v1.ListOrgReq) (*v1.ListOrgRes, error)
		// GetOrg 获取组织详情
		GetOrg(ctx context.Context, in *v1.GetOrgReq) (*v1.GetOrgRes, error)
		// CreateOrg 新建组织
		CreateOrg(ctx context.Context, in *v1.CreateOrgReq) (*v1.CreateOrgRes, error)
		// UpdateOrg 编辑组织
		UpdateOrg(ctx context.Context, in *v1.UpdateOrgReq) (*v1.UpdateOrgRes, error)
		// DeleteOrg 删除组织
		DeleteOrg(ctx context.Context, in *v1.DeleteOrgReq) (*v1.DeleteOrgRes, error)
		// MoveOrg 拖拽移动/排序
		MoveOrg(ctx context.Context, in *v1.MoveOrgReq) (*v1.MoveOrgRes, error)
	}
)

var (
	localOrg IOrg
)

func Org() IOrg {
	if localOrg == nil {
		panic("implement not found for interface IOrg, forgot register?")
	}
	return localOrg
}

func RegisterOrg(i IOrg) {
	localOrg = i
}
