// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure of table bbk_user_base for DAO operations like Where/Data.
type UserBase struct {
	g.Meta    `orm:"table:bbk_user_base, do:true"`
	Id        interface{} // 主键
	Group     interface{} // 用户组
	Password  interface{} // 密码
	LastTime  interface{} // 上次登陆时间
	LastIp    interface{} // 上次登陆ip
	IsDeleted interface{} // 数据状态0正常1删除
	CreateBy  interface{} // 创建人
	UpdateBy  interface{} // 修改人
	DeleteBy  interface{} // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
