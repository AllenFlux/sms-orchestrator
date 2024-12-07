// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalLeaseOrderDao is internal type for wrapping internal DAO implements.
type internalLeaseOrderDao = *internal.LeaseOrderDao

// leaseOrderDao is the data access object for table lease_order.
// You can define custom methods on it to extend its functionality as you wish.
type leaseOrderDao struct {
	internalLeaseOrderDao
}

var (
	// LeaseOrder is globally public accessible object for table lease_order operations.
	LeaseOrder = leaseOrderDao{
		internal.NewLeaseOrderDao(),
	}
)

// Fill with you ideas below.