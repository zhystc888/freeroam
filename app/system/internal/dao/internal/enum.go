// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EnumDao is the data access object for the table free_enum.
type EnumDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  EnumColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// EnumColumns defines and stores column names for the table free_enum.
type EnumColumns struct {
	Id        string // 主键
	Name      string // 枚举名称
	Code      string // 枚举编码
	Value     string // 枚举值
	ValueDesc string // 枚举值描述
	Module    string // 模块
	TableName string // 表名
	IsDeleted string // 数据状态0正常1删除
	CreateBy  string // 创建人
	UpdateBy  string // 修改人
	DeleteBy  string // 删除人
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// enumColumns holds the columns for the table free_enum.
var enumColumns = EnumColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Value:     "value",
	ValueDesc: "value_desc",
	Module:    "module",
	TableName: "table_name",
	IsDeleted: "is_deleted",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	DeleteBy:  "delete_by",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewEnumDao creates and returns a new DAO object for table data access.
func NewEnumDao(handlers ...gdb.ModelHandler) *EnumDao {
	return &EnumDao{
		group:    "default",
		table:    "free_enum",
		columns:  enumColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EnumDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EnumDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EnumDao) Columns() EnumColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EnumDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EnumDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EnumDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
