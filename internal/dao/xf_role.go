// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"SmartPower/internal/dao/internal"
)

// internalXfRoleDao is internal type for wrapping internal DAO implements.
type internalXfRoleDao = *internal.XfRoleDao

// xfRoleDao is the data access object for table xf_role.
// You can define custom methods on it to extend its functionality as you wish.
type xfRoleDao struct {
	internalXfRoleDao
}

var (
	// XfRole is globally public accessible object for table xf_role operations.
	XfRole = xfRoleDao{
		internal.NewXfRoleDao(),
	}
)

// Fill with you ideas below.
