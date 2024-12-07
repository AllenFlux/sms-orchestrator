// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalCollectTaskDao is internal type for wrapping internal DAO implements.
type internalCollectTaskDao = *internal.CollectTaskDao

// collectTaskDao is the data access object for table collect_task.
// You can define custom methods on it to extend its functionality as you wish.
type collectTaskDao struct {
	internalCollectTaskDao
}

var (
	// CollectTask is globally public accessible object for table collect_task operations.
	CollectTask = collectTaskDao{
		internal.NewCollectTaskDao(),
	}
)

// Fill with you ideas below.