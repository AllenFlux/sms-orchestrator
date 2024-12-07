// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalDeviceGroupDao is internal type for wrapping internal DAO implements.
type internalDeviceGroupDao = *internal.DeviceGroupDao

// deviceGroupDao is the data access object for table device_group.
// You can define custom methods on it to extend its functionality as you wish.
type deviceGroupDao struct {
	internalDeviceGroupDao
}

var (
	// DeviceGroup is globally public accessible object for table device_group operations.
	DeviceGroup = deviceGroupDao{
		internal.NewDeviceGroupDao(),
	}
)

// Fill with you ideas below.
