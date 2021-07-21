package controller

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
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
		return util.ResponseUnauthenticated("Tài khoản không tồn tại", err)
	}

	if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
		return util.ResponseUnauthenticated("Mật khẩu không chính xác", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserId
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	result, err := token.SignedString([]byte(config.GetJwt().Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return util.ResponseSuccess("Thành công", result)
}
