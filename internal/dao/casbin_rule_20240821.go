// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalCasbinRule20240821Dao is internal type for wrapping internal DAO implements.
type internalCasbinRule20240821Dao = *internal.CasbinRule20240821Dao

// casbinRule20240821Dao is the data access object for table casbin_rule_20240821.
// You can define custom methods on it to extend its functionality as you wish.
type casbinRule20240821Dao struct {
	internalCasbinRule20240821Dao
}

var (
	// CasbinRule20240821 is globally public accessible object for table casbin_rule_20240821 operations.
	CasbinRule20240821 = casbinRule20240821Dao{
		internal.NewCasbinRule20240821Dao(),
	}
)

// Fill with you ideas below.