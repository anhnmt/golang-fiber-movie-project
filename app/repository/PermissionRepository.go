package repository

import (
	"errors"
	model "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllPermissionsByStatus : Find permission by PermissionId and Status = 1
func FindAllPermissionsByStatus(status int) (*[]model.Permission, error) {
	permissions := make([]model.Permission, 0)

	if err := db.Find(&permissions, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &permissions, nil
}

// FindPermissionByIdAndStatus : Find permission by PermissionId and Status = 1
func FindPermissionByIdAndStatus(id string, status int) (*model.Permission, error) {
	uid := util.ParseStringToUInt(id)

	var permission model.Permission
	if err := db.Where(&model.Permission{PermissionId: uid, Status: status}).Find(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &permission, nil
}

func SavePermission(permission model.Permission) (*model.Permission, error) {
	if err := db.Save(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}
