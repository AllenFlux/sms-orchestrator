// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalMinerImportPhotoTaskDao is internal type for wrapping internal DAO implements.
type internalMinerImportPhotoTaskDao = *internal.MinerImportPhotoTaskDao

// minerImportPhotoTaskDao is the data access object for table miner_import_photo_task.
// You can define custom methods on it to extend its functionality as you wish.
type minerImportPhotoTaskDao struct {
	internalMinerImportPhotoTaskDao
}

var (
	// MinerImportPhotoTask is globally public accessible object for table miner_import_photo_task operations.
	MinerImportPhotoTask = minerImportPhotoTaskDao{
		internal.NewMinerImportPhotoTaskDao(),
	}
)

// Fill with you ideas below.
