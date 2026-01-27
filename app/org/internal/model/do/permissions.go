// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure of table free_permissions for DAO operations like Where/Data.
type Permissions struct {
	g.Meta    `orm:"table:free_permissions, do:true"`
	Id        any         // 主键ID
	ParentId  any         // 父权限ID（0为根）
	IdPath    any         // ID路径（/1/12/88/）
	CodePath  any         // Code路径（/org/org.member/org.member.list/）
	PermType  any         // 权限类型:permissions_type
	PermCode  any         // 权限标识
	Name      any         // 权限名称
	Desc      any         // 权限描述
	IsEnabled any         // 是否启用0:否,1:是
	Sort      any         // 排序
	IsDeleted any         // 是否删除0:否,1:是
	CreateBy  any         //
	UpdateBy  any         //
	DeleteBy  any         //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
	DeletedAt *gtime.Time //
}
