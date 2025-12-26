// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemEnumData is the golang structure of table free_system_enum_data for DAO operations like Where/Data.
type SystemEnumData struct {
	g.Meta        `orm:"table:free_system_enum_data, do:true"`
	Id            any         //
	EnumType      any         // 枚举类型，如 user_type
	EnumCode      any         // 枚举编码，如 ADMIN
	EnumValue     any         // 枚举值，如 1
	EnumLabel     any         // 前端展示文本
	EnumValueDesc any         // 枚举值说明
	Sort          any         // 顺序
	IsEnabled     any         // 是否启用0:否,1:是
	IsDeleted     any         // 是否删除0:否,1:是
	CreateBy      any         // 创建人
	UpdateBy      any         // 修改人
	DeleteBy      any         // 删除人
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
