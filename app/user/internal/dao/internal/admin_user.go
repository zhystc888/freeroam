// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserDao is the data access object for the table bbk_admin_user.
type AdminUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminUserColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminUserColumns defines and stores column names for the table bbk_admin_user.
type AdminUserColumns struct {
	Id                string // 主键
	UserId            string // 用户ID
	Username          string // 用户名
	Name              string // 名称
	ResetPasswordTime string // 重置密码时间
	Status            string // 状态：0未启用，1已启用，2禁止登陆
	Super             string // 超级管理员，0否1是
	IsDeleted         string // 数据状态0正常1删除
	CreateBy          string // 创建人
	UpdateBy          string // 修改人
	DeleteBy          string // 删除人
	CreateAt          string // 创建时间
	UpdateAt          string // 更新时间
	DeletedAt         string // 删除时间
}

// adminUserColumns holds the columns for the table bbk_admin_user.
var adminUserColumns = AdminUserColumns{
	Id:                "id",
	UserId:            "user_id",
	Username:          "username",
	Name:              "name",
	ResetPasswordTime: "reset_password_time",
	Status:            "status",
	Super:             "super",
	IsDeleted:         "is_deleted",
	CreateBy:          "create_by",
	UpdateBy:          "update_by",
	DeleteBy:          "delete_by",
	CreateAt:          "create_at",
	UpdateAt:          "update_at",
	DeletedAt:         "deleted_at",
}

// NewAdminUserDao creates and returns a new DAO object for table data access.
func NewAdminUserDao(handlers ...gdb.ModelHandler) *AdminUserDao {
	return &AdminUserDao{
		group:    "default",
		table:    "bbk_admin_user",
		columns:  adminUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminUserDao) Columns() AdminUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
