// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalEndpointDao is internal type for wrapping internal DAO implements.
type internalEndpointDao = *internal.EndpointDao

// endpointDao is the data access object for table endpoint.
// You can define custom methods on it to extend its functionality as you wish.
type endpointDao struct {
	internalEndpointDao
}

var (
	// Endpoint is globally public accessible object for table endpoint operations.
	Endpoint = endpointDao{
		internal.NewEndpointDao(),
	}
)

// Fill with you ideas below.