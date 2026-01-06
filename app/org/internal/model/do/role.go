// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table free_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:free_role, do:true"`
	Id        any         // 主键
	Code      any         // 角色编码(唯一)
	Name      any         // 角色名称
	Status    any         // 状态:role_status
	IsSystem  any         // 是否系统内置0:否,1:是
	Remark    any         // 备注
	IsDeleted any         // 是否删除0:否,1:是
	CreateBy  any         // 创建人
	UpdateBy  any         // 修改人
	DeleteBy  any         // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
