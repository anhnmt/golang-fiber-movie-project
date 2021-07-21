package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusOK

	return jsonResponse(c, status, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	status := fiber.StatusNotFound

	return jsonResponse(c, status, message, nil)
}

// ResponseError : returning json structure for error request
func ResponseError(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusInternalServerError

	return jsonResponse(c, status, message, data)
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusUnauthorized

	return jsonResponse(c, status, message, data)
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	status := fiber.StatusBadRequest

	return jsonResponse(c, status, message, data)
}

func jsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	if c != nil {
		ctx := c.Status(status)

		return ctx.JSON(&response.Response{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

	return fiberResponse(status, message, data)
}

func fiberResponse(status int, message string, data interface{}) error {
	msg, _ := JSONMarshal(&response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})

	return fiber.NewError(status, string(msg))
}
