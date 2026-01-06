// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`             // 主键
	Code      string      `json:"code"      orm:"code"       description:"角色编码(唯一)"`       // 角色编码(唯一)
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`           // 角色名称
	Status    string      `json:"status"    orm:"status"     description:"状态:role_status"` // 状态:role_status
	IsSystem  uint        `json:"isSystem"  orm:"is_system"  description:"是否系统内置0:否,1:是"`  // 是否系统内置0:否,1:是
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`             // 备注
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"是否删除0:否,1:是"`    // 是否删除0:否,1:是
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`            // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`            // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`            // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`           // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`           // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`           // 删除时间
}
