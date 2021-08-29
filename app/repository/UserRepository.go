package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

func FindAllUsersByStatus(status int) (*[]model.User, error) {
	users := make([]model.User, 0)

	err := db.Find(&users, "status = ?", status).Error

	return &users, err
}

// FindUserByUsernameAndStatus : Get User by Username and Status
func FindUserByUsernameAndStatus(username string, status int) (*model.User, error) {
	var user model.User

	err := db.
		Where("username = ? AND status = ?", username, status).
		Find(&user).Error

	return &user, err
}

func FindUserByIdAndStatus(id string, status int) (*model.User, error) {
	var user model.User

	err := db.
		Where("user_id = ? AND status = ?", id, status).
		Find(&user).Error

	return &user, err
}

func SaveUser(user model.User) (*model.User, error) {
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
