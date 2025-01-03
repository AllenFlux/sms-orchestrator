// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalRefundOrderDao is internal type for wrapping internal DAO implements.
type internalRefundOrderDao = *internal.RefundOrderDao

// refundOrderDao is the data access object for table refund_order.
// You can define custom methods on it to extend its functionality as you wish.
type refundOrderDao struct {
	internalRefundOrderDao
}

var (
	// RefundOrder is globally public accessible object for table refund_order operations.
	RefundOrder = refundOrderDao{
		internal.NewRefundOrderDao(),
	}
)

// Fill with you ideas below.
