// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrgDao is the data access object for the table free_org.
type OrgDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OrgColumns         // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OrgColumns defines and stores column names for the table free_org.
type OrgColumns struct {
	Id        string // 主键
	ParentId  string // 父组织ID, 0为顶级
	Name      string // 组织名称
	FullName  string // 组织全称
	Code      string // 组织编码
	Category  string // 分类:org_category
	Status    string // 状态:org_status
	Sort      string // 同级排序
	Path      string // 物化路径, 如 /1/10/
	IsDeleted string // 是否删除0:否,1:是
	CreateBy  string // 创建人
	UpdateBy  string // 修改人
	DeleteBy  string // 删除人
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// orgColumns holds the columns for the table free_org.
var orgColumns = OrgColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	FullName:  "full_name",
	Code:      "code",
	Category:  "category",
	Status:    "status",
	Sort:      "sort",
	Path:      "path",
	IsDeleted: "is_deleted",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	DeleteBy:  "delete_by",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewOrgDao creates and returns a new DAO object for table data access.
func NewOrgDao(handlers ...gdb.ModelHandler) *OrgDao {
	return &OrgDao{
		group:    "default",
		table:    "free_org",
		columns:  orgColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrgDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrgDao) Columns() OrgColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrgDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrgDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
