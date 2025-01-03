// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalSiteMessageReadDao is internal type for wrapping internal DAO implements.
type internalSiteMessageReadDao = *internal.SiteMessageReadDao

// siteMessageReadDao is the data access object for table site_message_read.
// You can define custom methods on it to extend its functionality as you wish.
type siteMessageReadDao struct {
	internalSiteMessageReadDao
}

var (
	// SiteMessageRead is globally public accessible object for table site_message_read operations.
	SiteMessageRead = siteMessageReadDao{
		internal.NewSiteMessageReadDao(),
	}
)

// Fill with you ideas below.
