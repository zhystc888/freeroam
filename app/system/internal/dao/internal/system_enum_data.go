// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemEnumDataDao is the data access object for the table free_system_enum_data.
type SystemEnumDataDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemEnumDataColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemEnumDataColumns defines and stores column names for the table free_system_enum_data.
type SystemEnumDataColumns struct {
	Id            string //
	EnumType      string // 枚举类型，如 user_type
	EnumCode      string // 枚举编码，如 ADMIN
	EnumValue     string // 枚举值，如 1
	EnumLabel     string // 前端展示文本
	EnumValueDesc string // 枚举值说明
	Sort          string // 顺序
	IsEnabled     string // 是否启用0:否,1:是
	IsDeleted     string // 是否删除0:否,1:是
	CreateBy      string // 创建人
	UpdateBy      string // 修改人
	DeleteBy      string // 删除人
	CreateAt      string // 创建时间
	UpdateAt      string // 更新时间
	DeletedAt     string // 删除时间
}

// systemEnumDataColumns holds the columns for the table free_system_enum_data.
var systemEnumDataColumns = SystemEnumDataColumns{
	Id:            "id",
	EnumType:      "enum_type",
	EnumCode:      "enum_code",
	EnumValue:     "enum_value",
	EnumLabel:     "enum_label",
	EnumValueDesc: "enum_value_desc",
	Sort:          "sort",
	IsEnabled:     "is_enabled",
	IsDeleted:     "is_deleted",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
	DeleteBy:      "delete_by",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
	DeletedAt:     "deleted_at",
}

// NewSystemEnumDataDao creates and returns a new DAO object for table data access.
func NewSystemEnumDataDao(handlers ...gdb.ModelHandler) *SystemEnumDataDao {
	return &SystemEnumDataDao{
		group:    "default",
		table:    "free_system_enum_data",
		columns:  systemEnumDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemEnumDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemEnumDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemEnumDataDao) Columns() SystemEnumDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemEnumDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemEnumDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemEnumDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
