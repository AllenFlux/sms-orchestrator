// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalSmsTraceTaskDao is internal type for wrapping internal DAO implements.
type internalSmsTraceTaskDao = *internal.SmsTraceTaskDao

// smsTraceTaskDao is the data access object for table sms_trace_task.
// You can define custom methods on it to extend its functionality as you wish.
type smsTraceTaskDao struct {
	internalSmsTraceTaskDao
}

var (
	// SmsTraceTask is globally public accessible object for table sms_trace_task operations.
	SmsTraceTask = smsTraceTaskDao{
		internal.NewSmsTraceTaskDao(),
	}
)

// Fill with you ideas below.
