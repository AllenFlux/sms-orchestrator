// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalInstallTaskItemDao is internal type for wrapping internal DAO implements.
type internalInstallTaskItemDao = *internal.InstallTaskItemDao

// installTaskItemDao is the data access object for table install_task_item.
// You can define custom methods on it to extend its functionality as you wish.
type installTaskItemDao struct {
	internalInstallTaskItemDao
}

var (
	// InstallTaskItem is globally public accessible object for table install_task_item operations.
	InstallTaskItem = installTaskItemDao{
		internal.NewInstallTaskItemDao(),
	}
)

// Fill with you ideas below.
