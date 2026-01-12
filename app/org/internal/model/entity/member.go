// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	Id               uint64      `json:"id"               orm:"id"                 description:"主键"`                // 主键
	Username         string      `json:"username"         orm:"username"           description:"账号(<=20), 不可重复"`    // 账号(<=20), 不可重复
	Name             string      `json:"name"             orm:"name"               description:"姓名(<=10)"`          // 姓名(<=10)
	Gender           uint        `json:"gender"           orm:"gender"             description:"性别:1男2女(可扩展0未知)"`   // 性别:1男2女(可扩展0未知)
	Mobile           string      `json:"mobile"           orm:"mobile"             description:"手机号(11位数字)"`        // 手机号(11位数字)
	Status           uint        `json:"status"           orm:"status"             description:"状态:1启用0禁用"`         // 状态:1启用0禁用
	IsSuperAdmin     uint        `json:"isSuperAdmin"     orm:"is_super_admin"     description:"超级管理员:1是0否"`        // 超级管理员:1是0否
	PasswordHash     string      `json:"passwordHash"     orm:"password_hash"      description:"密码hash"`            // 密码hash
	LastLoginAt      *gtime.Time `json:"lastLoginAt"      orm:"last_login_at"      description:"最近一次登录时间"`          // 最近一次登录时间
	ResignedAt       *gtime.Time `json:"resignedAt"       orm:"resigned_at"        description:"离职时间(非空表示离职)"`      // 离职时间(非空表示离职)
	SuperAdminUnique int         `json:"superAdminUnique" orm:"super_admin_unique" description:"仅用于保证超管唯一(业务字段勿用)"` // 仅用于保证超管唯一(业务字段勿用)
	IsDeleted        uint        `json:"isDeleted"        orm:"is_deleted"         description:"数据状态0正常1删除"`        // 数据状态0正常1删除
	CreateBy         uint64      `json:"createBy"         orm:"create_by"          description:"创建人"`               // 创建人
	UpdateBy         uint64      `json:"updateBy"         orm:"update_by"          description:"修改人"`               // 修改人
	DeleteBy         uint64      `json:"deleteBy"         orm:"delete_by"          description:"删除人"`               // 删除人
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"          description:"创建时间"`              // 创建时间
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"          description:"更新时间"`              // 更新时间
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`              // 删除时间
}
