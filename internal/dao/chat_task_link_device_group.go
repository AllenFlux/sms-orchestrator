// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalChatTaskLinkDeviceGroupDao is internal type for wrapping internal DAO implements.
type internalChatTaskLinkDeviceGroupDao = *internal.ChatTaskLinkDeviceGroupDao

// chatTaskLinkDeviceGroupDao is the data access object for table chat_task_link_device_group.
// You can define custom methods on it to extend its functionality as you wish.
type chatTaskLinkDeviceGroupDao struct {
	internalChatTaskLinkDeviceGroupDao
}

var (
	// ChatTaskLinkDeviceGroup is globally public accessible object for table chat_task_link_device_group operations.
	ChatTaskLinkDeviceGroup = chatTaskLinkDeviceGroupDao{
		internal.NewChatTaskLinkDeviceGroupDao(),
	}
)

// Fill with you ideas below.