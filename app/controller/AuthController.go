package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/response"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
	"github.com/xdorro/golang-fiber-movie-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"log"
	"sync"
	"time"
)

type AuthController struct {
	userRepository *repository.UserRepository
}

func NewAuthController() *AuthController {
	if authController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if authController == nil {
				authController = &AuthController{
					userRepository: repository.NewUserRepository(),
				}
				log.Println("Create new UserController")
			}
		})
	}

	return authController
}

// AuthToken : Find user by Username and Password and Status = 1
// @Summary Authentication User
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {object} dto.DataResponse{}
// @Failure 400 {object} dto.DataResponse{}
// @Router /api/oauth/token [post]
func (obj *AuthController) AuthToken(c *fiber.Ctx) error {
	var loginRequest request.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return util.ResponseBadRequest("Đăng nhập không thành công", err)
	}

	user, err := obj.userRepository.FindUserByUsernameAndStatus(loginRequest.Username, util.StatusActivated)
	if user == nil || user.Username == "" || err != nil {
		return util.ResponseBadRequest("Tài khoản không tồn tại", err)
	}

	if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
		return util.ResponseBadRequest("Mật khẩu không chính xác", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserId
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte(config.GetJwt().Secret))
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := &response.UserResponse{
		UserId:   user.UserId,
		Name:     user.Name,
		Username: user.Username,
		Gender:   user.Gender,
		Token:    accessToken,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *AuthController) CurrentUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	status := []int{util.StatusDraft, util.StatusDeleted}

	user, err := obj.userRepository.FindUserByUsernameAndStatusNotIn(username, status)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	result := mapper.UserSearch(user)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *AuthController) UpdateProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	status := []int{util.StatusDraft, util.StatusDeleted}

	user, err := obj.userRepository.FindUserByUsernameAndStatusNotIn(username, status)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	var profileRequest request.ProfileRequest

	if err = c.BodyParser(&profileRequest); err != nil {
		return util.ResponseBadRequest("Vui lòng nhập thông tin", err)
	}

	user.Name = profileRequest.Name
	user.Username = profileRequest.Username
	user.Gender = profileRequest.Gender

	if _, err = obj.userRepository.UpdateByUsername(username, *user); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *AuthController) ChangePassword(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	status := []int{util.StatusDraft, util.StatusDeleted}

	user, err := obj.userRepository.FindUserByUsernameAndStatusNotIn(username, status)

	if err != nil || user.UserId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	var passwordRequest request.ChangePasswordRequest

	if err = c.BodyParser(&passwordRequest); err != nil {
		return util.ResponseBadRequest("Vui lòng nhập thông tin", err)
	}

	if passwordRequest.OldPassword == "" || !util.CheckPasswordHash(passwordRequest.OldPassword, user.Password) {
		return util.ResponseUnauthorized("Mật khẩu cũ không chính xác", nil)
	}

	hash, _ := util.HashPassword(passwordRequest.NewPassword)
	user.Password = hash

	if _, err = obj.userRepository.UpdateByUsername(username, *user); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *AuthController) Restricted(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.SendString("Welcome " + username)
}
