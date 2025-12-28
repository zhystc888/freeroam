// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemConfig is the golang structure of table free_system_config for DAO operations like Where/Data.
type SystemConfig struct {
	g.Meta      `orm:"table:free_system_config, do:true"`
	Id          any         //
	ConfigCode  any         // 配置编码，如 token_validity_period
	ConfigValue any         // 配置值，如 60
	ConfigDesc  any         // 配置说明，如 token有效时长(分钟)
	IsDeleted   any         // 是否删除0:否,1:是
	CreateBy    any         // 创建人
	UpdateBy    any         // 修改人
	DeleteBy    any         // 删除人
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
