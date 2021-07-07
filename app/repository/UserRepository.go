package repository

import (
	"errors"
	model "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllUsersByStatus(status int) (*[]model.User, error) {
	users := make([]model.User, 0)

	if err := db.Find(&users, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &users, nil
}

// FindUserByUsernameAndStatus : Get User by Username and Status
func FindUserByUsernameAndStatus(username string, status int) (*model.User, error) {
	var user model.User

	if err := db.Where(&model.User{Username: username, Status: status}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func FindUserByIdAndStatus(id string, status int) (*model.User, error) {
	uid := util.ParseStringToUInt(id)

	var user model.User
	if err := db.Where(&model.User{UserId: uid, Status: status}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func SaveUser(user model.User) (*model.User, error) {
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
