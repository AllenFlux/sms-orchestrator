// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalTaskLinkDeviceGroupDao is internal type for wrapping internal DAO implements.
type internalTaskLinkDeviceGroupDao = *internal.TaskLinkDeviceGroupDao

// taskLinkDeviceGroupDao is the data access object for table task_link_device_group.
// You can define custom methods on it to extend its functionality as you wish.
type taskLinkDeviceGroupDao struct {
	internalTaskLinkDeviceGroupDao
}

var (
	// TaskLinkDeviceGroup is globally public accessible object for table task_link_device_group operations.
	TaskLinkDeviceGroup = taskLinkDeviceGroupDao{
		internal.NewTaskLinkDeviceGroupDao(),
	}
)

// Fill with you ideas below.