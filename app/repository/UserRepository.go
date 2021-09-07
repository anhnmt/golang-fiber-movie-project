package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if userRepository == nil {
				userRepository = &UserRepository{
					db: db,
				}
				log.Println("Create new UserRepository")
			}
		})
	}

	return userRepository
}

func (obj *UserRepository) FindAllUsersByStatus(status int) (*[]model.User, error) {
	users := make([]model.User, 0)

	err := obj.db.Find(&users, "status = ?", status).Error

	return &users, err
}

// FindUserByUsernameAndStatus : Get User by Username and Status
func (obj *UserRepository) FindUserByUsernameAndStatus(username string, status int) (*model.User, error) {
	var user model.User

	err := obj.db.
		Where("username = ? AND status = ?", username, status).
		Find(&user).Error

	return &user, err
}

func (obj *UserRepository) FindUserByIdAndStatus(id string, status int) (*model.User, error) {
	var user model.User

	err := obj.db.
		Where("user_id = ? AND status = ?", id, status).
		Find(&user).Error

	return &user, err
}

func (obj *UserRepository) FindUserByUsernameAndUserIdNotAndStatusNotIn(username string, id string, status []int) (*model.User, error) {
	var user model.User

	err := obj.db.
		Where("user_id <> ?", id).
		Where("username = ? AND status NOT IN ?", username, status).
		Find(&user).Error

	return &user, err
}

func (obj *UserRepository) FindUserByUsernameAndStatusNotIn(username string, status []int) (*model.User, error) {
	var user model.User

	err := obj.db.
		Where("username = ? AND status NOT IN ?", username, status).
		Find(&user).Error

	return &user, err
}

func (obj *UserRepository) SaveUser(user model.User) (*model.User, error) {
	if err := obj.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
