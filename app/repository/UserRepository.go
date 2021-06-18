package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindUserByUsernameAndStatus : Get User by Username and Status
func FindUserByUsernameAndStatus(username string, status int8) (*model.User, error) {
	var user model.User

	if err := db.Where(&model.User{Username: username, Status: status}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func FindUserByIdAndStatus(id string, status int8) (*model.User, error) {
	var user model.User

	uid := util.ParseStringToUInt(id)

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
