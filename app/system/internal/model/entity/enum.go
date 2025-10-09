// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Enum is the golang structure for table enum.
type Enum struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`         // 主键
	Name      string      `json:"name"      orm:"name"       description:"枚举名称"`       // 枚举名称
	Code      string      `json:"code"      orm:"code"       description:"枚举编码"`       // 枚举编码
	Value     int         `json:"value"     orm:"value"      description:"枚举值"`        // 枚举值
	ValueDesc string      `json:"valueDesc" orm:"value_desc" description:"枚举值描述"`      // 枚举值描述
	Module    string      `json:"module"    orm:"module"     description:"模块"`         // 模块
	TableName string      `json:"tableName" orm:"table_name" description:"表名"`         // 表名
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"数据状态0正常1删除"` // 数据状态0正常1删除
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:"创建人"`        // 创建人
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:"修改人"`        // 修改人
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:"删除人"`        // 删除人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`       // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`       // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`       // 删除时间
}
