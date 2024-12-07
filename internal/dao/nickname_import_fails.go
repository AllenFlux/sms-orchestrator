// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalNicknameImportFailsDao is internal type for wrapping internal DAO implements.
type internalNicknameImportFailsDao = *internal.NicknameImportFailsDao

// nicknameImportFailsDao is the data access object for table nickname_import_fails.
// You can define custom methods on it to extend its functionality as you wish.
type nicknameImportFailsDao struct {
	internalNicknameImportFailsDao
}

var (
	// NicknameImportFails is globally public accessible object for table nickname_import_fails operations.
	NicknameImportFails = nicknameImportFailsDao{
		internal.NewNicknameImportFailsDao(),
	}
)

// Fill with you ideas below.
