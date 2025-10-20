// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure of table free_user_base for DAO operations like Where/Data.
type UserBase struct {
	g.Meta    `orm:"table:free_user_base, do:true"`
	Id        any         // 主键
	TenantId  any         // 租户id
	Group     any         // 用户组
	Password  any         // 密码
	LastTime  any         // 上次登陆时间
	LastIp    any         // 上次登陆ip
	IsDeleted any         // 数据状态0正常1删除
	CreateBy  any         // 创建人
	UpdateBy  any         // 修改人
	DeleteBy  any         // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
