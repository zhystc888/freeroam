// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure for table user_base.
type UserBase struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`         // 主键
	Group     uint        `json:"group"     orm:"group"      description:"用户组"`        // 用户组
	Password  string      `json:"password"  orm:"password"   description:"密码"`         // 密码
	LastTime  uint64      `json:"lastTime"  orm:"last_time"  description:"上次登陆时间"`     // 上次登陆时间
	LastIp    string      `json:"lastIp"    orm:"last_ip"    description:"上次登陆ip"`     // 上次登陆ip
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"数据状态0正常1删除"` // 数据状态0正常1删除
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`        // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`        // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`        // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`       // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`       // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`       // 删除时间
}
