// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalPicTypeDao is internal type for wrapping internal DAO implements.
type internalPicTypeDao = *internal.PicTypeDao

// picTypeDao is the data access object for table pic_type.
// You can define custom methods on it to extend its functionality as you wish.
type picTypeDao struct {
	internalPicTypeDao
}

var (
	// PicType is globally public accessible object for table pic_type operations.
	PicType = picTypeDao{
		internal.NewPicTypeDao(),
	}
)

// Fill with you ideas below.