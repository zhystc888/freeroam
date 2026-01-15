// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure of table free_member for DAO operations like Where/Data.
type Member struct {
	g.Meta           `orm:"table:free_member, do:true"`
	Id               interface{} // 主键
	Username         interface{} // 账号(<=20), 不可重复
	Name             interface{} // 姓名(<=10)
	Gender           interface{} // 性别:1男2女(可扩展0未知)
	Mobile           interface{} // 手机号(11位数字)
	Status           interface{} // 状态:1启用0禁用
	IsSuperAdmin     interface{} // 超级管理员:1是0否
	PasswordHash     interface{} // 密码hash
	LastLoginAt      *gtime.Time // 最近一次登录时间
	ResignedAt       *gtime.Time // 离职时间(非空表示离职)
	SuperAdminUnique interface{} // 仅用于保证超管唯一(业务字段勿用)
	IsDeleted        interface{} // 数据状态0正常1删除
	CreateBy         interface{} // 创建人
	UpdateBy         interface{} // 修改人
	DeleteBy         interface{} // 删除人
	CreateAt         *gtime.Time // 创建时间
	UpdateAt         *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
