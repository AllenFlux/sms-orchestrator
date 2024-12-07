// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalMinerLoginTaskLinkDeviceDao is internal type for wrapping internal DAO implements.
type internalMinerLoginTaskLinkDeviceDao = *internal.MinerLoginTaskLinkDeviceDao

// minerLoginTaskLinkDeviceDao is the data access object for table miner_login_task_link_device.
// You can define custom methods on it to extend its functionality as you wish.
type minerLoginTaskLinkDeviceDao struct {
	internalMinerLoginTaskLinkDeviceDao
}

var (
	// MinerLoginTaskLinkDevice is globally public accessible object for table miner_login_task_link_device operations.
	MinerLoginTaskLinkDevice = minerLoginTaskLinkDeviceDao{
		internal.NewMinerLoginTaskLinkDeviceDao(),
	}
)

// Fill with you ideas below.
