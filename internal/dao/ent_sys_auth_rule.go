// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalEntSysAuthRuleDao is internal type for wrapping internal DAO implements.
type internalEntSysAuthRuleDao = *internal.EntSysAuthRuleDao

// entSysAuthRuleDao is the data access object for table ent_sys_auth_rule.
// You can define custom methods on it to extend its functionality as you wish.
type entSysAuthRuleDao struct {
	internalEntSysAuthRuleDao
}

var (
	// EntSysAuthRule is globally public accessible object for table ent_sys_auth_rule operations.
	EntSysAuthRule = entSysAuthRuleDao{
		internal.NewEntSysAuthRuleDao(),
	}
)

// Fill with you ideas below.
