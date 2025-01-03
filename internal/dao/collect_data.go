// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalCollectDataDao is internal type for wrapping internal DAO implements.
type internalCollectDataDao = *internal.CollectDataDao

// collectDataDao is the data access object for table collect_data.
// You can define custom methods on it to extend its functionality as you wish.
type collectDataDao struct {
	internalCollectDataDao
}

var (
	// CollectData is globally public accessible object for table collect_data operations.
	CollectData = collectDataDao{
		internal.NewCollectDataDao(),
	}
)

// Fill with you ideas below.
