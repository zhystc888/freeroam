// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrgStructure is the golang structure for table org_structure.
type OrgStructure struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`         // 主键
	TenantId  uint64      `json:"tenantId"  orm:"tenant_id"  description:"租户id"`       // 租户id
	Name      string      `json:"name"      orm:"name"       description:"组织名称"`       // 组织名称
	Level     uint        `json:"level"     orm:"level"      description:"组织级别"`       // 组织级别
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  description:"父id"`        // 父id
	Type      uint        `json:"type"      orm:"type"       description:"组织分类，枚举"`    // 组织分类，枚举
	Code      string      `json:"code"      orm:"code"       description:"组织编码"`       // 组织编码
	Status    uint        `json:"status"    orm:"status"     description:"组织状态，枚举"`    // 组织状态，枚举
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"数据状态0正常1删除"` // 数据状态0正常1删除
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`        // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`        // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`        // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`       // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`       // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`       // 删除时间
}
