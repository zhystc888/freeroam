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
	Id        interface{} // 主键
	TenantId  interface{} // 租户id
	Name      interface{} // 组织名称
	Level     interface{} // 组织级别
	ParentId  interface{} // 父id
	Type      interface{} // 组织分类，枚举
	Code      interface{} // 组织编码
	Status    interface{} // 组织状态，枚举
	IsDeleted interface{} // 数据状态0正常1删除
	CreateBy  interface{} // 创建人
	UpdateBy  interface{} // 修改人
	DeleteBy  interface{} // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
