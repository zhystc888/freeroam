// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberDao is the data access object for the table free_member.
type MemberDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MemberColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MemberColumns defines and stores column names for the table free_member.
type MemberColumns struct {
	Id               string // 主键
	Username         string // 账号(<=20), 不可重复
	Name             string // 姓名(<=10)
	Gender           string // 性别:1男2女(可扩展0未知)
	Mobile           string // 手机号(11位数字)
	Status           string // 状态:1启用0禁用
	IsSuperAdmin     string // 超级管理员:1是0否
	PasswordHash     string // 密码hash
	LastLoginAt      string // 最近一次登录时间
	ResignedAt       string // 离职时间(非空表示离职)
	SuperAdminUnique string // 仅用于保证超管唯一(业务字段勿用)
	IsDeleted        string // 数据状态0正常1删除
	CreateBy         string // 创建人
	UpdateBy         string // 修改人
	DeleteBy         string // 删除人
	CreateAt         string // 创建时间
	UpdateAt         string // 更新时间
	DeletedAt        string // 删除时间
}

// memberColumns holds the columns for the table free_member.
var memberColumns = MemberColumns{
	Id:               "id",
	Username:         "username",
	Name:             "name",
	Gender:           "gender",
	Mobile:           "mobile",
	Status:           "status",
	IsSuperAdmin:     "is_super_admin",
	PasswordHash:     "password_hash",
	LastLoginAt:      "last_login_at",
	ResignedAt:       "resigned_at",
	SuperAdminUnique: "super_admin_unique",
	IsDeleted:        "is_deleted",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
	DeleteBy:         "delete_by",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
	DeletedAt:        "deleted_at",
}

// NewMemberDao creates and returns a new DAO object for table data access.
func NewMemberDao(handlers ...gdb.ModelHandler) *MemberDao {
	return &MemberDao{
		group:    "default",
		table:    "free_member",
		columns:  memberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MemberDao) Columns() MemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
