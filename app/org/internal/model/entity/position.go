// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Position is the golang structure for table position.
type Position struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`                       // 主键
	Name      string      `json:"name"      orm:"name"       description:"职务名称"`                     // 职务名称
	Status    string      `json:"status"    orm:"status"     description:"状态:position_status"`       // 状态:position_status
	DataScope string      `json:"dataScope" orm:"data_scope" description:"数据权限:position_data_scope"` // 数据权限:position_data_scope
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"是否删除0:否,1:是"`              // 是否删除0:否,1:是
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`                      // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`                      // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`                      // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`                     // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`                     // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`                     // 删除时间
}
