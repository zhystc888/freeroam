// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemConfig is the golang structure for table system_config.
type SystemConfig struct {
	Id          int64       `json:"id"          orm:"id"           description:""`                             //
	ConfigCode  string      `json:"configCode"  orm:"config_code"  description:"配置编码，如 token_validity_period"` // 配置编码，如 token_validity_period
	ConfigValue string      `json:"configValue" orm:"config_value" description:"配置值，如 60"`                     // 配置值，如 60
	ConfigDesc  string      `json:"configDesc"  orm:"config_desc"  description:"配置说明，如 token有效时长(分钟)"`         // 配置说明，如 token有效时长(分钟)
	IsDeleted   uint        `json:"isDeleted"   orm:"is_deleted"   description:"是否删除0:否,1:是"`                  // 是否删除0:否,1:是
	CreateBy    uint64      `json:"createBy"    orm:"create_by"    description:"创建人"`                          // 创建人
	UpdateBy    uint64      `json:"updateBy"    orm:"update_by"    description:"修改人"`                          // 修改人
	DeleteBy    uint64      `json:"deleteBy"    orm:"delete_by"    description:"删除人"`                          // 删除人
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    description:"创建时间"`                         // 创建时间
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    description:"更新时间"`                         // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`                         // 删除时间
}
