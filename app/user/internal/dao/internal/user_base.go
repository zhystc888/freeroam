// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBaseDao is the data access object for the table bbk_user_base.
type UserBaseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserBaseColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserBaseColumns defines and stores column names for the table bbk_user_base.
type UserBaseColumns struct {
	Id        string // 主键
	Group     string // 用户组
	Password  string // 密码
	LastTime  string // 上次登陆时间
	LastIp    string // 上次登陆ip
	IsDeleted string // 数据状态0正常1删除
	CreateBy  string // 创建人
	UpdateBy  string // 修改人
	DeleteBy  string // 删除人
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// userBaseColumns holds the columns for the table bbk_user_base.
var userBaseColumns = UserBaseColumns{
	Id:        "id",
	Group:     "group",
	Password:  "password",
	LastTime:  "last_time",
	LastIp:    "last_ip",
	IsDeleted: "is_deleted",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	DeleteBy:  "delete_by",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewUserBaseDao creates and returns a new DAO object for table data access.
func NewUserBaseDao(handlers ...gdb.ModelHandler) *UserBaseDao {
	return &UserBaseDao{
		group:    "default",
		table:    "bbk_user_base",
		columns:  userBaseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserBaseDao) Columns() UserBaseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserBaseDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
