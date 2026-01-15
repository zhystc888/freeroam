package v1

import (
	"freeroam/common/model"

	"github.com/gogf/gf/v2/frame/g"
)

// ListPositionReq 职务列表请求
type ListPositionReq struct {
	g.Meta `path:"/org/positions" tags:"职务管理" method:"get" summary:"职务列表"`
	model.PageReq
	// 职务名称模糊
	Keyword string `p:"keyword" dc:"职务名称模糊"`
	// 状态:position_status
	Status string `p:"status" v:"enum:position_status" dc:"状态:position_status"`
}

// PositionListItem 职务列表项
type PositionListItem struct {
	// 职务 ID
	Id int64 `json:"id" dc:"职务ID"`
	// 职务名称
	Name string `json:"name" dc:"职务名称"`
	// 状态:position_status
	Status string `json:"status" dc:"状态:position_status"`
	// 数据权限:position_data_scope
	DataScope string `json:"dataScope" dc:"数据权限:position_data_scope"`
	// 关联的组织 ID数组
	OrgIds []int64 `json:"orgIds" dc:"关联的组织ID数组"`
	// 关联的角色 ID数组
	RoleIds []int64 `json:"roleIds" dc:"关联的角色ID数组"`
	// 创建时间
	CreateAt int64 `json:"createAt" dc:"创建时间"`
}

// ListPositionRes 职务列表响应
type ListPositionRes = model.ListRes[PositionListItem]

// GetPositionReq 获取职务详情请求
type GetPositionReq struct {
	g.Meta `path:"/org/positions/{id}" tags:"职务管理" method:"get" summary:"获取职务详情"`
	// 职务 ID
	Id int64 `p:"id" v:"required|min:1#职务ID不能为空|职务ID必须大于0" dc:"职务ID"`
}

// GetPositionRes 获取职务详情响应
type GetPositionRes struct {
	// 职务 ID
	Id int64 `json:"id" dc:"职务ID"`
	// 职务名称
	Name string `json:"name" dc:"职务名称"`
	// 状态:position_status
	Status string `json:"status" dc:"状态:position_status"`
	// 数据权限:position_data_scope
	DataScope string `json:"dataScope" dc:"数据权限:position_data_scope"`
	// 关联的组织 ID数组
	OrgIds []int64 `json:"orgIds" dc:"关联的组织ID数组"`
	// 关联的角色 ID数组
	RoleIds []int64 `json:"roleIds" dc:"关联的角色ID数组"`
	// 创建时间
	CreateAt int64 `json:"createAt" dc:"创建时间"`
	// 更新时间
	UpdateAt int64 `json:"updateAt" dc:"更新时间"`
}

// CreatePositionReq 新建职务请求
type CreatePositionReq struct {
	g.Meta `path:"/org/positions" tags:"职务管理" method:"post" summary:"新建职务"`
	// 职务名称
	Name string `json:"name" v:"required|length:1,20#职务名称不能为空|职务名称长度必须在1-20之间" dc:"职务名称"`
	// 状态:position_status
	Status string `json:"status" v:"required|enum:position_status#状态不能为空" dc:"状态:position_status"`
	// 数据权限:position_data_scope
	DataScope string `json:"dataScope" v:"required|enum:position_data_scope#数据权限不能为空" dc:"数据权限:position_data_scope"`
	// 关联的组织 ID数组
	OrgIds []int64 `json:"orgIds" dc:"关联的组织ID数组"`
	// 关联的角色 ID数组
	RoleIds []int64 `json:"roleIds" dc:"关联的角色ID数组"`
}

// CreatePositionRes 新建职务响应
type CreatePositionRes struct {
	// 职务 ID
	Id int64 `json:"id" dc:"职务ID"`
}

// UpdatePositionReq 编辑职务请求
type UpdatePositionReq struct {
	g.Meta `path:"/org/positions/{id}" tags:"职务管理" method:"put" summary:"编辑职务"`
	// 职务 ID
	Id int64 `p:"id" v:"required|min:1#职务ID不能为空|职务ID必须大于0" dc:"职务ID"`
	// 职务名称
	Name string `json:"name" v:"length:1,20#职务名称长度必须在1-20之间" dc:"职务名称"`
	// 状态:position_status
	Status string `json:"status" v:"enum:position_status" dc:"状态:position_status"`
	// 数据权限:position_data_scope
	DataScope string `json:"dataScope" v:"enum:position_data_scope" dc:"数据权限:position_data_scope"`
	// 关联的组织 ID数组
	OrgIds []int64 `json:"orgIds" dc:"关联的组织ID数组"`
	// 关联的角色 ID数组
	RoleIds []int64 `json:"roleIds" dc:"关联的角色ID数组"`
}

// UpdatePositionRes 编辑职务响应
type UpdatePositionRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// DeletePositionReq 删除职务请求
type DeletePositionReq struct {
	g.Meta `path:"/org/positions/{id}" tags:"职务管理" method:"delete" summary:"删除职务"`
	// 职务 ID
	Id int64 `p:"id" v:"required|min:1#职务ID不能为空|职务ID必须大于0" dc:"职务ID"`
}

// DeletePositionRes 删除职务响应
type DeletePositionRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// GetPositionOptionsReq 按组织获取可选职务集合请求
type GetPositionOptionsReq struct {
	g.Meta `path:"/org/positions/options" tags:"职务管理" method:"get" summary:"按组织获取可选职务集合"`
	// 组织ID列表(逗号分隔)
	OrgIds string `p:"orgIds" v:"required#组织ID不能为空" dc:"组织ID列表(逗号分隔)"`
	// 状态过滤,默认仅返回启用
	Status string `p:"status" v:"enum:position_status" dc:"状态过滤,默认仅返回启用"`
	// 职务名称模糊
	Keyword string `p:"keyword" dc:"职务名称模糊"`
}

// PositionOption 职务选项
type PositionOption struct {
	// 职务 ID
	Id int64 `json:"id" dc:"职务ID"`
	// 职务名称
	Name string `json:"name" dc:"职务名称"`
	// 状态:position_status
	Status string `json:"status" dc:"状态:position_status"`
	// 数据权限:position_data_scope
	DataScope string `json:"dataScope" dc:"数据权限:position_data_scope"`
}

// GetPositionOptionsRes 按组织获取可选职务集合响应
type GetPositionOptionsRes struct {
	// 职务选项列表
	List []*PositionOption `json:"list" dc:"职务选项列表"`
}
