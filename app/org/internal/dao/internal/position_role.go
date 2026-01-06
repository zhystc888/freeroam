// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionRoleDao is the data access object for the table free_position_role.
type PositionRoleDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PositionRoleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PositionRoleColumns defines and stores column names for the table free_position_role.
type PositionRoleColumns struct {
	Id         string // 主键
	PositionId string // 职务ID
	RoleId     string // 角色ID
	IsDeleted  string // 是否删除0:否,1:是
	CreateBy   string // 创建人
	UpdateBy   string // 修改人
	DeleteBy   string // 删除人
	CreateAt   string // 创建时间
	UpdateAt   string // 更新时间
	DeletedAt  string // 删除时间
}

// positionRoleColumns holds the columns for the table free_position_role.
var positionRoleColumns = PositionRoleColumns{
	Id:         "id",
	PositionId: "position_id",
	RoleId:     "role_id",
	IsDeleted:  "is_deleted",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
	DeleteBy:   "delete_by",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeletedAt:  "deleted_at",
}

// NewPositionRoleDao creates and returns a new DAO object for table data access.
func NewPositionRoleDao(handlers ...gdb.ModelHandler) *PositionRoleDao {
	return &PositionRoleDao{
		group:    "default",
		table:    "free_position_role",
		columns:  positionRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PositionRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PositionRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PositionRoleDao) Columns() PositionRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PositionRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PositionRoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PositionRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
