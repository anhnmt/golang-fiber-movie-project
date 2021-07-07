package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusOK

	return response(c, status, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	status := fiber.StatusNotFound

	return response(c, status, message, nil)
}

// ResponseError : returning json structure for error request
func ResponseError(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusInternalServerError

	return response(c, status, message, data)
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusUnauthorized

	return response(c, status, message, data)
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusBadRequest

	return response(c, status, message, data)
}

func response(c *fiber.Ctx, status int, message string, data interface{}) error {
	ctx := c.Status(status)

	if data != nil {
		return ctx.JSON(&request.DataResponse{
			DefaultResponse: request.DefaultResponse{
				Status:  status,
				Message: message,
			},
			Data: data,
		})
	}

	return ctx.JSON(&request.DefaultResponse{
		Status:  status,
		Message: message,
	})
}
