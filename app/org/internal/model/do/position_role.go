// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionRole is the golang structure of table free_position_role for DAO operations like Where/Data.
type PositionRole struct {
	g.Meta     `orm:"table:free_position_role, do:true"`
	Id         any         // 主键
	PositionId any         // 职务ID
	RoleId     any         // 角色ID
	IsDeleted  any         // 是否删除0:否,1:是
	CreateBy   any         // 创建人
	UpdateBy   any         // 修改人
	DeleteBy   any         // 删除人
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
