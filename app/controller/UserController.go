package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

// FindAllUsers : Find all users by Status = 1
func FindAllUsers(c *fiber.Ctx) error {
	db := database.GetDB()
	var users []model.User
	db.Find(&users, "status = ?", 1)
	return util.ResponseSuccess(c, "Thành công", users)
}

// FindUserById : Find user by User_Id and Status = 1
func FindUserById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	userId := c.Params("id")
	user := new(model.User)

	user, err = findUserByIdAndStatus(userId, user, db)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	return util.ResponseSuccess(c, "Thành công", user)
}

// CreateNewUser : Create a new user
func CreateNewUser(c *fiber.Ctx) error {
	db := database.GetDB()
	userRequest := new(dto.UserRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	user := model.User{
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: userRequest.Password,
		Gender:   userRequest.Gender,
	}

	if err := db.Create(&user).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateUserById : Update user by User_Id and Status = 1
func UpdateUserById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	userId := c.Params("id")
	userRequest := new(dto.UserRequest)
	user := new(model.User)

	user, err = findUserByIdAndStatus(userId, user, db)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if err := c.BodyParser(userRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	user.Name = userRequest.Name
	user.Username = userRequest.Username
	user.Password = userRequest.Password
	user.Gender = userRequest.Gender

	if err := db.Save(&user).Error; err != nil {
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

	user, err = findUserByIdAndStatus(userId, user, db)

	if err != nil || user.UserId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if result := db.Model(&user).Update("status", 0); result.Error != nil {
		return util.ResponseError(c, result.Error.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// findUserByIdAndStatus : Find user by User_Id and Status = 1
func findUserByIdAndStatus(userId string, user *model.User, db *gorm.DB) (*model.User, error) {
	if err := db.First(&user, "user_id = ? and status = ?", userId, 1).Error; err != nil {
		return nil, err
	}

	return user, nil
}
