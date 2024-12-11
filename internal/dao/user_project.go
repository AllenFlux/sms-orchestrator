// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalUserProjectDao is internal type for wrapping internal DAO implements.
type internalUserProjectDao = *internal.UserProjectDao

// userProjectDao is the data access object for table user_project.
// You can define custom methods on it to extend its functionality as you wish.
type userProjectDao struct {
	internalUserProjectDao
}

var (
	// UserProject is globally public accessible object for table user_project operations.
	UserProject = userProjectDao{
		internal.NewUserProjectDao(),
	}
)

// Fill with you ideas below.
