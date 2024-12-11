// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalLogDao is internal type for wrapping internal DAO implements.
type internalLogDao = *internal.LogDao

// logDao is the data access object for table log.
// You can define custom methods on it to extend its functionality as you wish.
type logDao struct {
	internalLogDao
}

var (
	// Log is globally public accessible object for table log operations.
	Log = logDao{
		internal.NewLogDao(),
	}
)

// Fill with you ideas below.
