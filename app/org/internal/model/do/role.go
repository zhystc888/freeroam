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
	Id        interface{} // 主键
	TenantId  interface{} // 租户id
	RoleName  interface{} // 角色名称
	Level     interface{} // 角色级别
	Disabled  interface{} // 是否禁用0否1是
	IsDeleted interface{} // 数据状态0正常1删除
	CreateBy  interface{} // 创建人
	UpdateBy  interface{} // 修改人
	DeleteBy  interface{} // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
