package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateRoleReq 创建角色请求
type CreateRoleReq struct {
	g.Meta `path:"/org/role" tags:"角色管理" method:"post" summary:"创建角色"`
	// 角色编码(唯一)
	Code string `json:"code" v:"required|length:1,64#角色编码不能为空|角色编码长度必须在1-64之间" dc:"角色编码"`
	// 角色名称
	Name string `json:"name" v:"required|length:1,64#角色名称不能为空|角色名称长度必须在1-64之间" dc:"角色名称"`
	// 状态:role_status
	Status string `json:"status" v:"required|enum:role_status#状态不能为空" dc:"状态"`
	// 备注
	Remark string `json:"remark" dc:"备注"`
}

// CreateRoleRes 创建角色响应
type CreateRoleRes struct {
	// 角色ID
	Id int64 `json:"id" dc:"角色ID"`
}

// UpdateRoleReq 更新角色请求
type UpdateRoleReq struct {
	g.Meta `path:"/org/role" tags:"角色管理" method:"put" summary:"更新角色"`
	// 角色ID
	Id int64 `json:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
	// 角色名称
	Name string `json:"name" v:"length:1,64#角色名称长度必须在1-64之间" dc:"角色名称"`
	// 状态:role_status
	Status string `json:"status" v:"enum:role_status" dc:"状态"`
	// 备注
	Remark string `json:"remark" dc:"备注"`
}

// UpdateRoleRes 更新角色响应
type UpdateRoleRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// DeleteRoleReq 删除角色请求
type DeleteRoleReq struct {
	g.Meta `path:"/org/role" tags:"角色管理" method:"delete" summary:"删除角色"`
	// 角色ID
	Id int64 `p:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
}

// DeleteRoleRes 删除角色响应
type DeleteRoleRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// GetRoleReq 获取角色详情请求
type GetRoleReq struct {
	g.Meta `path:"/org/role/{id}" tags:"角色管理" method:"get" summary:"获取角色详情"`
	// 角色ID
	Id int64 `p:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
}

// GetRoleRes 获取角色详情响应
type GetRoleRes struct {
	// 角色ID
	Id int64 `json:"id" dc:"角色ID"`
	// 角色编码(唯一)
	Code string `json:"code" dc:"角色编码"`
	// 角色名称
	Name string `json:"name" dc:"角色名称"`
	// 状态:role_status
	Status string `json:"status" dc:"状态"`
	// 是否系统内置0:否,1:是
	IsSystem bool `json:"isSystem" dc:"是否系统内置"`
	// 备注
	Remark string `json:"remark" dc:"备注"`
	// 创建时间
	CreateAt int64 `json:"createAt" dc:"创建时间"`
	// 更新时间
	UpdateAt int64 `json:"updateAt" dc:"更新时间"`
}

// ListRoleReq 获取角色列表请求
type ListRoleReq struct {
	g.Meta `path:"/org/role" tags:"角色管理" method:"get" summary:"获取角色列表"`
	// 页码
	Page int64 `p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	// 每页数量
	PageSize int64 `p:"pageSize" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
	// 对 code/name 模糊
	Keyword string `p:"keyword" dc:"对 code/name 模糊"`
	// 状态:role_status
	Status string `p:"status" dc:"状态"`
	// 是否系统内置0:否,1:是
	IsSystem int64 `p:"isSystem" dc:"是否系统内置"`
}

// ListRoleRes 获取角色列表响应
type ListRoleRes struct {
	// 角色列表
	List []*GetRoleRes `json:"list" dc:"角色列表"`
	// 总数
	Total int64 `json:"total" dc:"总数"`
	// 页码
	Page int64 `json:"page" dc:"页码"`
	// 每页数量
	PageSize int64 `json:"pageSize" dc:"每页数量"`
}

// GetRolePositionListReq 查询角色绑定的职务列表请求
type GetRolePositionListReq struct {
	g.Meta `path:"/org/role/{roleId}/positions" tags:"角色管理" method:"get" summary:"查询角色绑定的职务列表"`
	// 角色ID
	RoleId int64 `p:"roleId" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
	// 页码
	Page int64 `p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	// 每页数量
	PageSize int64 `p:"pageSize" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
	// 对 code/name 模糊
	Keyword string `p:"keyword" dc:"对 code/name 模糊"`
}

// GetRolePositionListRes 查询角色绑定的职务列表响应
type GetRolePositionListRes struct {
	// 职务列表
	List []*PositionItem `json:"list" dc:"职务列表"`
	// 总数
	Total int64 `json:"total" dc:"总数"`
	// 页码
	Page int64 `json:"page" dc:"页码"`
	// 每页数量
	PageSize int64 `json:"pageSize" dc:"每页数量"`
}

// PositionItem 职务项
type PositionItem struct {
	// 职务ID
	PositionId int64 `json:"positionId" dc:"职务ID"`
	// 职务名称
	PositionName string `json:"positionName" dc:"职务名称"`
	// 职务状态
	PositionStatus string `json:"positionStatus" dc:"职务状态"`
	// 绑定时间
	CreateAt string `json:"createAt" dc:"绑定时间"`
}

// BatchAssignRolePositionReq 批量绑定职务到角色（覆盖式）请求
type BatchAssignRolePositionReq struct {
	g.Meta `path:"/org/role/{roleId}/positions" tags:"角色管理" method:"post" summary:"批量绑定职务到角色（覆盖式）"`
	// 角色ID
	RoleId int64 `p:"roleId" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
	// 职务ID列表
	PositionIds []int64 `json:"positionIds" dc:"职务ID列表"`
}

// BatchAssignRolePositionRes 批量绑定职务到角色（覆盖式）响应
type BatchAssignRolePositionRes struct {
	// 是否成功
	Success bool `json:"success" dc:"是否成功"`
}

// GetRolePositionIdsReq 查询角色当前绑定职务ID集合（表单回显）请求
type GetRolePositionIdsReq struct {
	g.Meta `path:"/org/role/{roleId}/position-ids" tags:"角色管理" method:"get" summary:"查询角色当前绑定职务ID集合（表单回显）"`
	// 角色ID
	RoleId int64 `p:"roleId" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
}

// GetRolePositionIdsRes 查询角色当前绑定职务ID集合（表单回显）响应
type GetRolePositionIdsRes struct {
	// 职务ID列表
	PositionIds []int64 `json:"positionIds" dc:"职务ID列表"`
}
