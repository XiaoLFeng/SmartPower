// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"SmartPower/internal/dao/internal"
)

// internalXfUserDao is internal type for wrapping internal DAO implements.
type internalXfUserDao = *internal.XfUserDao

// xfUserDao is the data access object for table xf_user.
// You can define custom methods on it to extend its functionality as you wish.
type xfUserDao struct {
	internalXfUserDao
}

var (
	// XfUser is globally public accessible object for table xf_user operations.
	XfUser = xfUserDao{
		internal.NewXfUserDao(),
	}
)

// Fill with you ideas below.