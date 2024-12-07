// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalApplyListDao is internal type for wrapping internal DAO implements.
type internalApplyListDao = *internal.ApplyListDao

// applyListDao is the data access object for table apply_list.
// You can define custom methods on it to extend its functionality as you wish.
type applyListDao struct {
	internalApplyListDao
}

var (
	// ApplyList is globally public accessible object for table apply_list operations.
	ApplyList = applyListDao{
		internal.NewApplyListDao(),
	}
)

// Fill with you ideas below.