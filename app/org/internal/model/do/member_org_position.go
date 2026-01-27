// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberOrgPosition is the golang structure of table free_member_org_position for DAO operations like Where/Data.
type MemberOrgPosition struct {
	g.Meta     `orm:"table:free_member_org_position, do:true"`
	Id         any         // 主键
	MemberId   any         // 成员ID
	OrgUnitId  any         // 组织ID
	PositionId any         // 职务ID
	IsDeleted  any         // 数据状态0正常1删除
	CreateBy   any         // 创建人
	UpdateBy   any         // 修改人
	DeleteBy   any         // 删除人
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
