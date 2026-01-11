// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"context"

	"freeroam/app/gateway/api/org/v1"
)

type IOrgV1 interface {
	GetOrgTree(ctx context.Context, req *v1.GetOrgTreeReq) (res *v1.GetOrgTreeRes, err error)
	ListOrg(ctx context.Context, req *v1.ListOrgReq) (res *v1.ListOrgRes, err error)
	GetOrg(ctx context.Context, req *v1.GetOrgReq) (res *v1.GetOrgRes, err error)
	CreateOrg(ctx context.Context, req *v1.CreateOrgReq) (res *v1.CreateOrgRes, err error)
	UpdateOrg(ctx context.Context, req *v1.UpdateOrgReq) (res *v1.UpdateOrgRes, err error)
	DeleteOrg(ctx context.Context, req *v1.DeleteOrgReq) (res *v1.DeleteOrgRes, err error)
	MoveOrg(ctx context.Context, req *v1.MoveOrgReq) (res *v1.MoveOrgRes, err error)
}
