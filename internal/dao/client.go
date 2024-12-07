// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalClientDao is internal type for wrapping internal DAO implements.
type internalClientDao = *internal.ClientDao

// clientDao is the data access object for table client.
// You can define custom methods on it to extend its functionality as you wish.
type clientDao struct {
	internalClientDao
}

var (
	// Client is globally public accessible object for table client operations.
	Client = clientDao{
		internal.NewClientDao(),
	}
)

// Fill with you ideas below.
