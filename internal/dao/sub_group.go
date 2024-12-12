// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalSubGroupDao is internal type for wrapping internal DAO implements.
type internalSubGroupDao = *internal.SubGroupDao

// subGroupDao is the data access object for table sub_group.
// You can define custom methods on it to extend its functionality as you wish.
type subGroupDao struct {
	internalSubGroupDao
}

var (
	// SubGroup is globally public accessible object for table sub_group operations.
	SubGroup = subGroupDao{
		internal.NewSubGroupDao(),
	}
)

// Fill with you ideas below.