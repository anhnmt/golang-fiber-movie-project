package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllUserRolesByUserIdAndStatus(id string, status int8) (*[]model.Role, error) {
	uid := util.ParseStringToUInt(id)

	roles := make([]model.Role, 0)
	if err := db.
		Model(&model.User{}).
		Select("roles.*").
		Joins("LEFT JOIN user_roles on users.user_id = user_roles.user_id").
		Joins("LEFT JOIN roles on roles.role_id = user_roles.role_id").
		Where(&model.User{UserId: uid, Status: status}).
		Find(&roles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &roles, nil
}
