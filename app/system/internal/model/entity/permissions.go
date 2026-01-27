// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure for table permissions.
type Permissions struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`                                     // 主键ID
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  description:"父权限ID（0为根）"`                               // 父权限ID（0为根）
	IdPath    string      `json:"idPath"    orm:"id_path"    description:"ID路径（/1/12/88/）"`                          // ID路径（/1/12/88/）
	CodePath  string      `json:"codePath"  orm:"code_path"  description:"Code路径（/org/org.member/org.member.list/）"` // Code路径（/org/org.member/org.member.list/）
	PermType  uint        `json:"permType"  orm:"perm_type"  description:"权限类型:permissions_type"`                    // 权限类型:permissions_type
	PermCode  string      `json:"permCode"  orm:"perm_code"  description:"权限标识"`                                     // 权限标识
	Name      string      `json:"name"      orm:"name"       description:"权限名称"`                                     // 权限名称
	Desc      string      `json:"desc"      orm:"desc"       description:"权限描述"`                                     // 权限描述
	IsEnabled uint        `json:"isEnabled" orm:"is_enabled" description:"是否启用0:否,1:是"`                              // 是否启用0:否,1:是
	Sort      uint        `json:"sort"      orm:"sort"       description:"排序"`                                       // 排序
	IsDeleted uint        `json:"isDeleted" orm:"is_deleted" description:"是否删除0:否,1:是"`                              // 是否删除0:否,1:是
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  description:""`                                         //
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  description:""`                                         //
	DeleteBy  uint64      `json:"deleteBy"  orm:"delete_by"  description:""`                                         //
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:""`                                         //
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:""`                                         //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`                                         //
}
