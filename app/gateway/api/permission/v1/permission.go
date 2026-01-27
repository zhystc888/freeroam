package v1

import "github.com/gogf/gf/v2/frame/g"

// AssignRolePermissionsReq 角色权限分配请求
type AssignRolePermissionsReq struct {
	g.Meta `path:"/org/permissions/{roleId}" tags:"权限管理" method:"put" summary:"角色权限分配"`
	// 角色ID
	RoleId int64 `p:"roleId" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
	// 权限标识列表（页面/组件）
	PermCodes []string `json:"permCodes" v:"required|min-length:1#权限标识列表不能为空|至少需要一个权限标识" dc:"权限标识列表"`
}

// AssignRolePermissionsRes 角色权限分配响应
type AssignRolePermissionsRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// GetRolePermissionsReq 查询角色已分配权限请求
type GetRolePermissionsReq struct {
	g.Meta `path:"/org/permissions/{roleId}" tags:"权限管理" method:"get" summary:"查询角色已分配权限"`
	// 角色ID
	RoleId int64 `p:"roleId" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
}

// GetRolePermissionsRes 查询角色已分配权限响应
type GetRolePermissionsRes struct {
	// 权限 ID列表
	PermissionIds []int64 `json:"permissionIds" dc:"权限ID列表"`
}

// GetPermissionTreeReq 查询权限资源树请求
type GetPermissionTreeReq struct {
	g.Meta `path:"/org/permissions/tree" tags:"权限管理" method:"get" summary:"查询权限资源树" perm:"org:permissions:tree:get"`
}

// GetPermissionTreeRes 查询权限资源树响应
type GetPermissionTreeRes struct {
	// 权限树
	Tree []*PermissionTreeNode `json:"tree" dc:"权限树"`
}

// PermissionTreeNode 权限树节点
type PermissionTreeNode struct {
	// 权限 ID
	Id int64 `json:"id" dc:"权限ID"`
	// 权限标识
	PermCode string `json:"permCode" dc:"权限标识"`
	// 权限名称
	Name string `json:"name" dc:"权限名称"`
	// 权限类型:permissions_type
	PermType string `json:"permType" dc:"权限类型:permissions_type"`
	// 子节点
	Children []*PermissionTreeNode `json:"children" dc:"子节点"`
}

// GetFrontPermissionsReq 获取前端权限集合请求
type GetFrontPermissionsReq struct {
	g.Meta `path:"/authz/front-permissions" tags:"权限管理" method:"get" summary:"获取前端权限集合"`
}

// GetFrontPermissionsRes 获取前端权限集合响应
type GetFrontPermissionsRes struct {
	// 权限标识列表（页面+组件）
	Permissions []string `json:"permissions" dc:"权限标识列表"`
}
