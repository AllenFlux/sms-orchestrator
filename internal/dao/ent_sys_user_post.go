// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalEntSysUserPostDao is internal type for wrapping internal DAO implements.
type internalEntSysUserPostDao = *internal.EntSysUserPostDao

// entSysUserPostDao is the data access object for table ent_sys_user_post.
// You can define custom methods on it to extend its functionality as you wish.
type entSysUserPostDao struct {
	internalEntSysUserPostDao
}

var (
	// EntSysUserPost is globally public accessible object for table ent_sys_user_post operations.
	EntSysUserPost = entSysUserPostDao{
		internal.NewEntSysUserPostDao(),
	}
)

// Fill with you ideas below.
