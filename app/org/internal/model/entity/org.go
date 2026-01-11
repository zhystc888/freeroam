// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Org is the golang structure for table org.
type Org struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`              // 主键
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  description:"父组织ID, 0为顶级"`     // 父组织ID, 0为顶级
	Name      string      `json:"name"      orm:"name"       description:"组织名称"`            // 组织名称
	FullName  string      `json:"fullName"  orm:"full_name"  description:"组织全称"`            // 组织全称
	Code      string      `json:"code"      orm:"code"       description:"组织编码"`            // 组织编码
	Category  string      `json:"category"  orm:"category"   description:"分类:org_category"` // 分类:org_category
	Status    string      `json:"status"    orm:"status"     description:"状态:org_status"`   // 状态:org_status
	Sort      int         `json:"sort"      orm:"sort"       description:"同级排序"`            // 同级排序
	Path      string      `json:"path"      orm:"path"       description:"物化路径, 如 /1/10/"`  // 物化路径, 如 /1/10/
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"是否删除0:否,1:是"`     // 是否删除0:否,1:是
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`             // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`             // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`             // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`            // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`            // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`            // 删除时间
}
