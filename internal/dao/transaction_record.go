// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalTransactionRecordDao is internal type for wrapping internal DAO implements.
type internalTransactionRecordDao = *internal.TransactionRecordDao

// transactionRecordDao is the data access object for table transaction_record.
// You can define custom methods on it to extend its functionality as you wish.
type transactionRecordDao struct {
	internalTransactionRecordDao
}

var (
	// TransactionRecord is globally public accessible object for table transaction_record operations.
	TransactionRecord = transactionRecordDao{
		internal.NewTransactionRecordDao(),
	}
)

// Fill with you ideas below.