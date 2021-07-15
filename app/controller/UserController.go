package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllUsers : Find all users by Status = 1
func FindAllUsers(c *fiber.Ctx) error {
	users, err := repository.FindAllUsersByStatus(util.StatusActivated)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	result := mapper.ListUserSearch(*users)

	return util.ResponseSuccess(c, "Thành công", result)
}

// FindUserById : Find user by User_Id and Status = 1
func FindUserById(c *fiber.Ctx) error {
	userId := c.Params("id")
	user, err := repository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	result := mapper.UserSearch(user)

	return util.ResponseSuccess(c, "Thành công", result)
}

// CreateNewUser : Create a new user
func CreateNewUser(c *fiber.Ctx) error {
	userRequest := new(request.UserRequest)

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
		Status:   util.StatusActivated,
	}

	if _, err = repository.SaveUser(user); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateUserById : Update user by User_Id and Status = 1
func UpdateUserById(c *fiber.Ctx) error {
	userId := c.Params("id")

	user, err := repository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	userRequest := new(request.UserRequest)
	if err = c.BodyParser(userRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	hash, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return util.ResponseError(c, "Không thể mã hoá mật khẩu", err)
	}

	user.Name = userRequest.Name
	user.Username = userRequest.Username
	user.Password = hash
	user.Gender = userRequest.Gender

	if _, err = repository.SaveUser(*user); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteUserById : Delete user by User_Id and Status = 1
func DeleteUserById(c *fiber.Ctx) error {
	userId := c.Params("id")

	user, err := repository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	user.Status = util.StatusDeleted

	if _, err = repository.SaveUser(*user); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
