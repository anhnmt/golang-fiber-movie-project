package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetJwt().Secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return util.ResponseUnauthorized("Token bị thiếu hoặc không đúng định dạng", nil)
	}

	return util.ResponseUnauthorized("Token không hợp lệ hoặc hết hạn", nil)
}
