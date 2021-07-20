package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusOK

	return customResponse(c, status, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	status := fiber.StatusNotFound

	return customResponse(c, status, message, nil)
}

// ResponseError : returning json structure for error request
func ResponseError(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusInternalServerError

	return customResponse(c, status, message, data)
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusUnauthorized

	return customResponse(c, status, message, data)
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusBadRequest

	return customResponse(c, status, message, data)
}

func customResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	ctx := c.Status(status)

	if data != nil {
		return ctx.JSON(&response.DataResponse{
			DefaultResponse: response.DefaultResponse{
				Status:  status,
				Message: message,
			},
			Data: data,
		})
	}

	return ctx.JSON(&response.DefaultResponse{
		Status:  status,
		Message: message,
	})
}
