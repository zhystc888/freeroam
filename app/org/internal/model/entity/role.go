// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`         // 主键
	TenantId  uint64      `json:"tenantId"  orm:"tenant_id"  description:"租户id"`       // 租户id
	RoleName  string      `json:"roleName"  orm:"role_name"  description:"角色名称"`       // 角色名称
	Level     uint        `json:"level"     orm:"level"      description:"角色级别"`       // 角色级别
	Disabled  uint        `json:"disabled"  orm:"disabled"   description:"是否禁用0否1是"`   // 是否禁用0否1是
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"数据状态0正常1删除"` // 数据状态0正常1删除
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`        // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`        // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`        // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`       // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`       // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`       // 删除时间
}
