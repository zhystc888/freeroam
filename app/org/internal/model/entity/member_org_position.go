// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberOrgPosition is the golang structure for table member_org_position.
type MemberOrgPosition struct {
	Id         uint64      `json:"id"         orm:"id"          description:"主键"`         // 主键
	MemberId   uint64      `json:"memberId"   orm:"member_id"   description:"成员ID"`       // 成员ID
	OrgUnitId  uint64      `json:"orgUnitId"  orm:"org_unit_id" description:"组织ID"`       // 组织ID
	PositionId uint64      `json:"positionId" orm:"position_id" description:"职务ID"`       // 职务ID
	IsDeleted  uint        `json:"isDeleted"  orm:"is_deleted"  description:"数据状态0正常1删除"` // 数据状态0正常1删除
	CreateBy   uint64      `json:"createBy"   orm:"create_by"   description:"创建人"`        // 创建人
	UpdateBy   uint64      `json:"updateBy"   orm:"update_by"   description:"修改人"`        // 修改人
	DeleteBy   uint64      `json:"deleteBy"   orm:"delete_by"   description:"删除人"`        // 删除人
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:"创建时间"`       // 创建时间
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:"更新时间"`       // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除时间"`       // 删除时间
}
