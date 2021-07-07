package repository

import (
	"errors"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllUsersByStatus(status int) (*[]model2.User, error) {
	users := make([]model2.User, 0)

	if err := db.Find(&users, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &users, nil
}

// FindUserByUsernameAndStatus : Get User by Username and Status
func FindUserByUsernameAndStatus(username string, status int) (*model2.User, error) {
	var user model2.User

	if err := db.Where(&model2.User{Username: username, Status: status}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func FindUserByIdAndStatus(id string, status int) (*model2.User, error) {
	uid := util.ParseStringToUInt(id)

	var user model2.User
	if err := db.Where(&model2.User{UserId: uid, Status: status}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func SaveUser(user model2.User) (*model2.User, error) {
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
