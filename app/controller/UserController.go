package controller

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

type UserController struct {
	userRepository *repository.UserRepository
}

func NewUserController() *UserController {
	if userController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if userController == nil {
				userController = &UserController{
					userRepository: repository.NewUserRepository(),
				}
				log.Println("Create new UserController")
			}
		})
	}

	return userController
}

// FindAllUsers : Find all users by Status = 1
func (obj *UserController) FindAllUsers(c *fiber.Ctx) error {
	users, err := obj.userRepository.FindAllUsersByStatus(util.StatusActivated)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.ListUserSearch(*users)

	return util.ResponseSuccess("Thành công", result)
}

// FindUserById : Find user by User_Id and Status = 1
func (obj *UserController) FindUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user, err := obj.userRepository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	result := mapper.UserSearch(user)

	return util.ResponseSuccess("Thành công", result)
}

// CreateNewUser : Create a new user
func (obj *UserController) CreateNewUser(c *fiber.Ctx) error {
	userRequest := new(request.UserRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	hash, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return util.ResponseError("Không thể mã hoá mật khẩu", err)
	}

	user := model.User{
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: hash,
		Gender:   userRequest.Gender,
		Status:   util.StatusActivated,
	}

	if _, err = obj.userRepository.SaveUser(user); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateUserById : Update user by User_Id and Status = 1
func (obj *UserController) UpdateUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	user, err := obj.userRepository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	userRequest := new(request.UserRequest)
	if err = c.BodyParser(userRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	user.Name = userRequest.Name
	user.Username = userRequest.Username
	user.Gender = userRequest.Gender

	if _, err = obj.userRepository.UpdateUser(userId, *user); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteUserById : Delete user by User_Id and Status = 1
func (obj *UserController) DeleteUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	user, err := obj.userRepository.FindUserByIdAndStatus(userId, util.StatusActivated)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	user.Status = util.StatusDeleted

	if _, err = obj.userRepository.UpdateUser(userId, *user); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *UserController) CheckIsExistUsername(c *fiber.Ctx) error {
	var err error
	var user *model.User

	username := c.Query("username")
	userId := c.Query("user_id")
	status := []int{util.StatusDraft, util.StatusDeleted}

	if userId != "" {
		user, err = obj.userRepository.FindUserByUsernameAndUserIdNotAndStatusNotIn(username, userId, status)
	} else {
		user, err = obj.userRepository.FindUserByUsernameAndStatusNotIn(username, status)
	}

	if err != nil || user.UserId == 0 {
		return util.ResponseSuccess("Thành công", false)
	}

	return util.ResponseSuccess("Thành công", true)
}
