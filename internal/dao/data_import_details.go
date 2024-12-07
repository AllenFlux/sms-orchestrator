// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalDataImportDetailsDao is internal type for wrapping internal DAO implements.
type internalDataImportDetailsDao = *internal.DataImportDetailsDao

// dataImportDetailsDao is the data access object for table data_import_details.
// You can define custom methods on it to extend its functionality as you wish.
type dataImportDetailsDao struct {
	internalDataImportDetailsDao
}

var (
	// DataImportDetails is globally public accessible object for table data_import_details operations.
	DataImportDetails = dataImportDetailsDao{
		internal.NewDataImportDetailsDao(),
	}
)

// Fill with you ideas below.