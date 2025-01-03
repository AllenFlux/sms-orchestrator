// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalMinerTaskLinkDeviceDao is internal type for wrapping internal DAO implements.
type internalMinerTaskLinkDeviceDao = *internal.MinerTaskLinkDeviceDao

// minerTaskLinkDeviceDao is the data access object for table miner_task_link_device.
// You can define custom methods on it to extend its functionality as you wish.
type minerTaskLinkDeviceDao struct {
	internalMinerTaskLinkDeviceDao
}

var (
	// MinerTaskLinkDevice is globally public accessible object for table miner_task_link_device operations.
	MinerTaskLinkDevice = minerTaskLinkDeviceDao{
		internal.NewMinerTaskLinkDeviceDao(),
	}
)

// Fill with you ideas below.
