// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalChatTaskDao is internal type for wrapping internal DAO implements.
type internalChatTaskDao = *internal.ChatTaskDao

// chatTaskDao is the data access object for table chat_task.
// You can define custom methods on it to extend its functionality as you wish.
type chatTaskDao struct {
	internalChatTaskDao
}

var (
	// ChatTask is globally public accessible object for table chat_task operations.
	ChatTask = chatTaskDao{
		internal.NewChatTaskDao(),
	}
)

// Fill with you ideas below.