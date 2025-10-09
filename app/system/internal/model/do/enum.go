// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Enum is the golang structure of table free_enum for DAO operations like Where/Data.
type Enum struct {
	g.Meta    `orm:"table:free_enum, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 枚举名称
	Code      interface{} // 枚举编码
	Value     interface{} // 枚举值
	ValueDesc interface{} // 枚举值描述
	Module    interface{} // 模块
	TableName interface{} // 表名
	IsDeleted interface{} // 数据状态0正常1删除
	CreateBy  interface{} // 创建人
	UpdateBy  interface{} // 修改人
	DeleteBy  interface{} // 删除人
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
