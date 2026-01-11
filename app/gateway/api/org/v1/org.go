package v1

import "github.com/gogf/gf/v2/frame/g"

// GetOrgTreeReq 获取组织树请求
type GetOrgTreeReq struct {
	g.Meta `path:"/org/units/tree" tags:"组织管理" method:"get" summary:"获取组织树"`
	// 组织名称模糊搜索
	Keyword string `p:"keyword" dc:"组织名称模糊搜索"`
}

// OrgTreeNode 组织树节点
type OrgTreeNode struct {
	// 组织 ID
	Id int64 `json:"id" dc:"组织ID"`
	// 组织名称
	Name string `json:"name" dc:"组织名称"`
	// 组织全称
	FullName string `json:"fullName" dc:"组织全称"`
	// 组织编码
	Code string `json:"code" dc:"组织编码"`
	// 分类:org_category
	Category string `json:"category" dc:" 分类:org_category"`
	// 状态:org_status
	Status string `json:"status" dc:"状态:org_status"`
	// 该组织成员数量
	MemberCount int64 `json:"memberCount" dc:"该组织成员数量"`
	// 子组织数组
	Children []*OrgTreeNode `json:"children" dc:"子组织数组"`
}

// GetOrgTreeRes 获取组织树响应
type GetOrgTreeRes struct {
	// 组织树列表
	List []*OrgTreeNode `json:"list" dc:"组织树列表"`
}

// ListOrgReq 获取组织列表请求
type ListOrgReq struct {
	g.Meta `path:"/org/units" tags:"组织管理" method:"get" summary:"获取组织列表"`
	// 页码
	Page int64 `p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	// 每页数量
	PageSize int64 `p:"pageSize" v:"min:1|max:200#每页数量必须大于0|每页数量不能超过200" dc:"每页数量"`
	// 组织名称模糊
	Keyword string `p:"keyword" dc:"组织名称模糊"`
	// 组织编码精确
	Code string `p:"code" dc:"组织编码精确"`
	// 分类:org_category
	Category string `p:"category" v:"enum:org_category" dc:" 分类:org_category"`
	// 状态:org_status
	Status string `p:"status" v:"enum:org_status" dc:"状态:org_status"`
}

// OrgListItem 组织列表项
type OrgListItem struct {
	// 组织 ID
	Id int64 `json:"id" dc:"组织ID"`
	// 组织名称
	Name string `json:"name" dc:"组织名称"`
	// 组织全称
	FullName string `json:"fullName" dc:"组织全称"`
	// 组织编码
	Code string `json:"code" dc:"组织编码"`
	// 状态:org_status
	Status string `json:"status" dc:"状态:org_status"`
	// 分类:org_category
	Category string `json:"category" dc:" 分类:org_category"`
	// 创建时间
	CreateAt int64 `json:"createAt" dc:"创建时间"`
}

// ListOrgRes 获取组织列表响应
type ListOrgRes struct {
	// 组织列表
	List []*OrgListItem `json:"list" dc:"组织列表"`
	// 总数
	Total int64 `json:"total" dc:"总数"`
	// 页码
	Page int64 `json:"page" dc:"页码"`
	// 每页数量
	PageSize int64 `json:"pageSize" dc:"每页数量"`
}

// GetOrgReq 获取组织详情请求
type GetOrgReq struct {
	g.Meta `path:"/org/units/{id}" tags:"组织管理" method:"get" summary:"获取组织详情"`
	// 组织 ID
	Id int64 `p:"id" v:"required|min:1#组织ID不能为空|组织ID必须大于0" dc:"组织ID"`
}

// GetOrgRes 获取组织详情响应
type GetOrgRes struct {
	// 组织 ID
	Id int64 `json:"id" dc:"组织ID"`
	// 父组织 ID
	ParentId int64 `json:"parentId" dc:"父组织ID"`
	// 组织名称
	Name string `json:"name" dc:"组织名称"`
	// 组织全称
	FullName string `json:"fullName" dc:"组织全称"`
	// 组织编码
	Code string `json:"code" dc:"组织编码"`
	// 分类:org_category
	Category string `json:"category" dc:" 分类:org_category"`
	// 状态:org_status
	Status string `json:"status" dc:"状态:org_status"`
	// 同级排序
	Sort int32 `json:"sort" dc:"同级排序"`
	// 物化路径
	Path string `json:"path" dc:"物化路径"`
	// 创建时间
	CreateAt int64 `json:"createAt" dc:"创建时间"`
	// 更新时间
	UpdateAt int64 `json:"updateAt" dc:"更新时间"`
}

// CreateOrgReq 新建组织请求
type CreateOrgReq struct {
	g.Meta `path:"/org/units" tags:"组织管理" method:"post" summary:"新建组织"`
	// 上级组织ID（0表示顶级）
	ParentId int64 `json:"parentId" dc:"上级组织ID（0表示顶级）"`
	// 组织名称
	Name string `json:"name" v:"required|length:1,30#组织名称不能为空|组织名称长度必须在1-30之间" dc:"组织名称"`
	// 组织全称
	FullName string `json:"fullName" dc:"组织全称"`
	// 组织编码
	Code string `json:"code" v:"required|length:1,20#组织编码不能为空|组织编码长度必须在1-20之间" dc:"组织编码"`
	// 分类:org_category
	Category string `json:"category" v:"required|enum:org_category#分类不能为空" dc:" 分类:org_category"`
	// 状态:org_status
	Status string `json:"status" v:"required|enum:org_status#状态不能为空" dc:"状态:org_status"`
	// 同级排序
	Sort int32 `json:"sort" dc:"同级排序"`
}

// CreateOrgRes 新建组织响应
type CreateOrgRes struct {
	// 组织 ID
	Id int64 `json:"id" dc:"组织ID"`
}

// UpdateOrgReq 编辑组织请求
type UpdateOrgReq struct {
	g.Meta `path:"/org/units/{id}" tags:"组织管理" method:"put" summary:"编辑组织"`
	// 组织 ID
	Id int64 `p:"id" v:"required|min:1#组织ID不能为空|组织ID必须大于0" dc:"组织ID"`
	// 上级组织 ID
	ParentId int64 `json:"parentId" dc:"上级组织ID"`
	// 组织名称
	Name string `json:"name" v:"length:1,30#组织名称长度必须在1-30之间" dc:"组织名称"`
	// 组织全称
	FullName string `json:"fullName" dc:"组织全称"`
	// 组织编码
	Code string `json:"code" v:"length:1,20#组织编码长度必须在1-20之间" dc:"组织编码"`
	// 分类:org_category
	Category string `json:"category" v:"enum:org_category" dc:" 分类:org_category"`
	// 状态:org_status
	Status string `json:"status" v:"enum:org_status" dc:"状态:org_status"`
	// 同级排序
	Sort int32 `json:"sort" dc:"同级排序"`
}

// UpdateOrgRes 编辑组织响应
type UpdateOrgRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// DeleteOrgReq 删除组织请求
type DeleteOrgReq struct {
	g.Meta `path:"/org/units/{id}" tags:"组织管理" method:"delete" summary:"删除组织"`
	// 组织 ID
	Id int64 `p:"id" v:"required|min:1#组织ID不能为空|组织ID必须大于0" dc:"组织ID"`
}

// DeleteOrgRes 删除组织响应
type DeleteOrgRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// MoveOrgReq 拖拽移动/排序请求
type MoveOrgReq struct {
	g.Meta `path:"/org/units/{id}/move" tags:"组织管理" method:"post" summary:"拖拽移动/排序"`
	// 组织 ID
	Id int64 `p:"id" v:"required|min:1#组织ID不能为空|组织ID必须大于0" dc:"组织ID"`
	// 新的父组织ID（0为顶级）
	NewParentId int64 `json:"newParentId" dc:"新的父组织ID（0为顶级）"`
	// 新的同级排序值
	NewSort int32 `json:"newSort" dc:"新的同级排序值"`
}

// MoveOrgRes 拖拽移动/排序响应
type MoveOrgRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}
