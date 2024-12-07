// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasbinRuleDao is the data access object for table casbin_rule.
type CasbinRuleDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CasbinRuleColumns // columns contains all the column names of Table for convenient usage.
}

// CasbinRuleColumns defines and stores column names for table casbin_rule.
type CasbinRuleColumns struct {
	Ptype string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

// casbinRuleColumns holds the columns for table casbin_rule.
var casbinRuleColumns = CasbinRuleColumns{
	Ptype: "ptype",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewCasbinRuleDao creates and returns a new DAO object for table data access.
func NewCasbinRuleDao() *CasbinRuleDao {
	return &CasbinRuleDao{
		group:   "default",
		table:   "casbin_rule",
		columns: casbinRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CasbinRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CasbinRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CasbinRuleDao) Columns() CasbinRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CasbinRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CasbinRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CasbinRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
