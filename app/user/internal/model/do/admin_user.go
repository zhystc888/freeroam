// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure of table bbk_admin_user for DAO operations like Where/Data.
type AdminUser struct {
	g.Meta            `orm:"table:bbk_admin_user, do:true"`
	Id                interface{} // 主键
	UserId            interface{} // 用户ID
	Username          interface{} // 用户名
	Name              interface{} // 名称
	ResetPasswordTime *gtime.Time // 重置密码时间
	Status            interface{} // 状态：0未启用，1已启用，2禁止登陆
	Super             interface{} // 超级管理员，0否1是
	IsDeleted         interface{} // 数据状态0正常1删除
	CreateBy          interface{} // 创建人
	UpdateBy          interface{} // 修改人
	DeleteBy          interface{} // 删除人
	CreateAt          *gtime.Time // 创建时间
	UpdateAt          *gtime.Time // 更新时间
	DeletedAt         *gtime.Time // 删除时间
}
