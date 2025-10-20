// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrgStructure is the golang structure of table free_org_structure for DAO operations like Where/Data.
type OrgStructure struct {
	g.Meta    `orm:"table:free_org_structure, do:true"`
	Id        any         // 主键
	TenantId  any         // 租户id
	Name      any         // 组织名称
	Level     any         // 组织级别
	ParentId  any         // 父id
	Type      any         // 组织分类，枚举
	Code      any         // 组织编码
	Status    any         // 组织状态，枚举
	IsDeleted any         // 数据状态0正常1删除
	CreateBy  any         // 创建人
	UpdateBy  any         // 修改人
	DeleteBy  any         // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
