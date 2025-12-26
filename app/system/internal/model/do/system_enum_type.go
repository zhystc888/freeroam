// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemEnumType is the golang structure of table free_system_enum_type for DAO operations like Where/Data.
type SystemEnumType struct {
	g.Meta       `orm:"table:free_system_enum_type, do:true"`
	Id           any         //
	EnumType     any         // 枚举类型，如 user_type
	EnumTypeDesc any         // 枚举类型说明
	IsEnabled    any         // 是否启用0:否,1:是
	IsDeleted    any         // 是否删除0:否,1:是
	CreateBy     any         // 创建人
	UpdateBy     any         // 修改人
	DeleteBy     any         // 删除人
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
