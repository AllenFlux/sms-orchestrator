// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalClientLinkLabelDao is internal type for wrapping internal DAO implements.
type internalClientLinkLabelDao = *internal.ClientLinkLabelDao

// clientLinkLabelDao is the data access object for table client_link_label.
// You can define custom methods on it to extend its functionality as you wish.
type clientLinkLabelDao struct {
	internalClientLinkLabelDao
}

var (
	// ClientLinkLabel is globally public accessible object for table client_link_label operations.
	ClientLinkLabel = clientLinkLabelDao{
		internal.NewClientLinkLabelDao(),
	}
)

// Fill with you ideas below.
