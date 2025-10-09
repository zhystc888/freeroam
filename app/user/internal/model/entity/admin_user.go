// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure for table admin_user.
type AdminUser struct {
	Id                uint64      `json:"id"                orm:"id"                  description:"主键"`                 // 主键
	UserId            uint64      `json:"userId"            orm:"user_id"             description:"用户ID"`               // 用户ID
	Username          string      `json:"username"          orm:"username"            description:"用户名"`                // 用户名
	Name              string      `json:"name"              orm:"name"                description:"名称"`                 // 名称
	ResetPasswordTime *gtime.Time `json:"resetPasswordTime" orm:"reset_password_time" description:"重置密码时间"`             // 重置密码时间
	Status            uint        `json:"status"            orm:"status"              description:"状态：0未启用，1已启用，2禁止登陆"` // 状态：0未启用，1已启用，2禁止登陆
	Super             uint        `json:"super"             orm:"super"               description:"超级管理员，0否1是"`         // 超级管理员，0否1是
	IsDeleted         uint        `json:"isDeleted"         orm:"is_deleted"          description:"数据状态0正常1删除"`         // 数据状态0正常1删除
	CreateBy          uint64      `json:"createBy"          orm:"create_by"           description:"创建人"`                // 创建人
	UpdateBy          uint64      `json:"updateBy"          orm:"update_by"           description:"修改人"`                // 修改人
	DeleteBy          uint64      `json:"deleteBy"          orm:"delete_by"           description:"删除人"`                // 删除人
	CreateAt          *gtime.Time `json:"createAt"          orm:"create_at"           description:"创建时间"`               // 创建时间
	UpdateAt          *gtime.Time `json:"updateAt"          orm:"update_at"           description:"更新时间"`               // 更新时间
	DeletedAt         *gtime.Time `json:"deletedAt"         orm:"deleted_at"          description:"删除时间"`               // 删除时间
}
