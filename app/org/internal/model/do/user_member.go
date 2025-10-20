// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMember is the golang structure of table free_user_member for DAO operations like Where/Data.
type UserMember struct {
	g.Meta            `orm:"table:free_user_member, do:true"`
	Id                any         // 主键
	UserId            any         // 用户ID
	TenantId          any         // 租户id
	Username          any         // 用户名
	Name              any         // 姓名
	Mobile            any         // 手机号
	Gender            any         // 性别
	ResetPasswordTime *gtime.Time // 重置密码时间
	Status            any         // 状态：0未启用，1已启用，2禁止登陆
	Super             any         // 超级管理员，0否1是
	IsDeleted         any         // 数据状态0正常1删除
	CreateBy          any         // 创建人
	UpdateBy          any         // 修改人
	DeleteBy          any         // 删除人
	CreateAt          *gtime.Time // 创建时间
	UpdateAt          *gtime.Time // 更新时间
	DeletedAt         *gtime.Time // 删除时间
}
