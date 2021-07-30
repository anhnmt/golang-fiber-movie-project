package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllPermissionsByStatus : Find permission by PermissionId and Status = 1
func FindAllPermissionsByStatus(status int) (*[]model.Permission, error) {
	permissions := make([]model.Permission, 0)

	err := db.Find(&permissions, "status = ?", status).Error

	return &permissions, err
}

// FindPermissionByIdAndStatus : Find permission by PermissionId and Status = 1
func FindPermissionByIdAndStatus(id string, status int) (*model.Permission, error) {
	uid := util.ParseStringToUInt(id)

	var permission model.Permission
	err := db.Where(&model.Permission{PermissionId: uid, Status: status}).Find(&permission).Error

	return &permission, err
}

func SavePermission(permission model.Permission) (*model.Permission, error) {
	err := db.Save(&permission).Error

	return &permission, err
}
