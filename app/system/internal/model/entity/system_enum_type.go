// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemEnumType is the golang structure for table system_enum_type.
type SystemEnumType struct {
	Id           int64       `json:"id"           orm:"id"             description:""`                 //
	EnumType     string      `json:"enumType"     orm:"enum_type"      description:"枚举类型，如 user_type"` // 枚举类型，如 user_type
	EnumTypeDesc string      `json:"enumTypeDesc" orm:"enum_type_desc" description:"枚举类型说明"`           // 枚举类型说明
	IsEnabled    int         `json:"isEnabled"    orm:"is_enabled"     description:"是否启用0:否,1:是"`      // 是否启用0:否,1:是
	IsDeleted    uint        `json:"isDeleted"    orm:"is_deleted"     description:"是否删除0:否,1:是"`      // 是否删除0:否,1:是
	CreateBy     uint64      `json:"createBy"     orm:"create_by"      description:"创建人"`              // 创建人
	UpdateBy     uint64      `json:"updateBy"     orm:"update_by"      description:"修改人"`              // 修改人
	DeleteBy     uint64      `json:"deleteBy"     orm:"delete_by"      description:"删除人"`              // 删除人
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"      description:"创建时间"`             // 创建时间
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"      description:"更新时间"`             // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"     description:"删除时间"`             // 删除时间
}
