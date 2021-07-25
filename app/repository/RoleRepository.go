package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

//func FindAllUserRolesByUserIdAndStatus(id string, status int) (*[]model.Role, error) {
//	uid := util.ParseStringToUInt(id)
//
//	roles := make([]model.Role, 0)
//	if err := db.
//		Model(&model.User{}).
//		Select("roles.*").
//		Joins("LEFT JOIN user_roles on users.user_id = user_roles.user_id").
//		Joins("LEFT JOIN roles on roles.role_id = user_roles.role_id").
//		Where(&model.User{UserId: uid, Status: status}).
//		Find(&roles).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil
//		}
//
//		return nil, err
//	}
//
//	return &roles, nil
//}

// FindAllRolesByStatus : Find role by RoleId and Status = 1
func FindAllRolesByStatus(status int) (*[]model.Role, error) {
	roles := make([]model.Role, 0)

	err := db.Find(&roles, "status = ?", status).Error

	return &roles, err
}

// FindRoleByIdAndStatus : Find role by RoleId and Status = 1
func FindRoleByIdAndStatus(id string, status int) (*model.Role, error) {
	uid := util.ParseStringToUInt(id)

	var role model.Role
	err := db.Where(&model.Role{RoleId: uid, Status: status}).Find(&role).Error

	return &role, err
}

func SaveRole(role model.Role) (*model.Role, error) {
	err := db.Save(&role).Error

	return &role, err
}
