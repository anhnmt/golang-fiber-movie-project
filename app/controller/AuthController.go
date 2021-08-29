package controller

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/response"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"time"
)

// AuthToken : Find user by Username and Password and Status = 1
// @Summary Authentication User
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {object} dto.DataResponse{}
// @Failure 400 {object} dto.DataResponse{}
// @Router /api/oauth/token [post]
func AuthToken(c *fiber.Ctx) error {
	var loginRequest request.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return util.ResponseBadRequest("Đăng nhập không thành công", err)
	}

	user, err := repository.FindUserByUsernameAndStatus(loginRequest.Username, util.StatusActivated)
	if user == nil || user.Username == "" || err != nil {
		return util.ResponseUnauthorized("Tài khoản không tồn tại", err)
	}

	if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
		return util.ResponseUnauthorized("Mật khẩu không chính xác", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

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
