package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

// FindAllUsers : Find all users by Status = 1
func FindAllUsers(c *fiber.Ctx) error {
	db := database.GetDB()
	var users []model.User
	db.Find(&users, "status = ?", 1)

	result := mapper.ListUserSearch(users)

	return util.ResponseSuccess(c, "Thành công", result)
}

// FindUserById : Find user by User_Id and Status = 1
func FindUserById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	userId := c.Params("id")
	user, err := findUserByIdAndStatus(userId, db)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	result := mapper.UserSearch(user)

	return util.ResponseSuccess(c, "Thành công", result)
}

// CreateNewUser : Create a new user
func CreateNewUser(c *fiber.Ctx) error {
	userRequest := new(dto.UserRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	hash, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return util.ResponseError(c, "Không thể mã hoá mật khẩu", err)
	}

	user := model.User{
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: hash,
		Gender:   userRequest.Gender,
	}

	if _, err = repository.SaveUser(user); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateUserById : Update user by User_Id and Status = 1
func UpdateUserById(c *fiber.Ctx) error {
	var err error

	userId := c.Params("id")
	userRequest := new(dto.UserRequest)
	var user *model.User

	user, err = repository.FindUserByIdAndStatus(userId, 1)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if err = c.BodyParser(userRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	hash, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return util.ResponseError(c, "Không thể mã hoá mật khẩu", err)
	}

	newUser := model.User{
		UserId:   user.UserId,
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: hash,
		Gender:   userRequest.Gender,
	}

	if _, err = repository.SaveUser(newUser); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteUserById : Delete user by User_Id and Status = 1
func DeleteUserById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	user := new(model.User)

	userId := c.Params("id")

	user, err = findUserByIdAndStatus(userId, db)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if result := db.Model(&user).Update("status", 0); result.Error != nil {
		return util.ResponseError(c, result.Error.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// findUserByIdAndStatus : Find user by User_Id and Status = 1
func findUserByIdAndStatus(userId string, db *gorm.DB) (*model.User, error) {
	user := new(model.User)

	if err := db.First(&user, "user_id = ? and status = ?", userId, 1).Error; err != nil {
		return nil, err
	}

	return user, nil
}
