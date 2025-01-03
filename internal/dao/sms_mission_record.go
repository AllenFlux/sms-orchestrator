// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalSmsMissionRecordDao is internal type for wrapping internal DAO implements.
type internalSmsMissionRecordDao = *internal.SmsMissionRecordDao

// smsMissionRecordDao is the data access object for table sms_mission_record.
// You can define custom methods on it to extend its functionality as you wish.
type smsMissionRecordDao struct {
	internalSmsMissionRecordDao
}

var (
	// SmsMissionRecord is globally public accessible object for table sms_mission_record operations.
	SmsMissionRecord = smsMissionRecordDao{
		internal.NewSmsMissionRecordDao(),
	}
)

// Fill with you ideas below.
