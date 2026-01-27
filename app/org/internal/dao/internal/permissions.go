// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionsDao is the data access object for the table free_permissions.
type PermissionsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PermissionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PermissionsColumns defines and stores column names for the table free_permissions.
type PermissionsColumns struct {
	Id        string // 主键ID
	ParentId  string // 父权限ID（0为根）
	IdPath    string // ID路径（/1/12/88/）
	CodePath  string // Code路径（/org/org.member/org.member.list/）
	PermType  string // 权限类型:permissions_type
	PermCode  string // 权限标识
	Name      string // 权限名称
	Desc      string // 权限描述
	IsEnabled string // 是否启用0:否,1:是
	Sort      string // 排序
	IsDeleted string // 是否删除0:否,1:是
	CreateBy  string //
	UpdateBy  string //
	DeleteBy  string //
	CreateAt  string //
	UpdateAt  string //
	DeletedAt string //
}

// permissionsColumns holds the columns for the table free_permissions.
var permissionsColumns = PermissionsColumns{
	Id:        "id",
	ParentId:  "parent_id",
	IdPath:    "id_path",
	CodePath:  "code_path",
	PermType:  "perm_type",
	PermCode:  "perm_code",
	Name:      "name",
	Desc:      "desc",
	IsEnabled: "is_enabled",
	Sort:      "sort",
	IsDeleted: "is_deleted",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	DeleteBy:  "delete_by",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewPermissionsDao creates and returns a new DAO object for table data access.
func NewPermissionsDao(handlers ...gdb.ModelHandler) *PermissionsDao {
	return &PermissionsDao{
		group:    "default",
		table:    "free_permissions",
		columns:  permissionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PermissionsDao) Columns() PermissionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
