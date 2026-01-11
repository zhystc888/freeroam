// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Org is the golang structure of table free_org for DAO operations like Where/Data.
type Org struct {
	g.Meta    `orm:"table:free_org, do:true"`
	Id        any         // 主键
	ParentId  any         // 父组织ID, 0为顶级
	Name      any         // 组织名称
	FullName  any         // 组织全称
	Code      any         // 组织编码
	Category  any         // 分类:org_category
	Status    any         // 状态:org_status
	Sort      any         // 同级排序
	Path      any         // 物化路径, 如 /1/10/
	IsDeleted any         // 是否删除0:否,1:是
	CreateBy  any         // 创建人
	UpdateBy  any         // 修改人
	DeleteBy  any         // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
