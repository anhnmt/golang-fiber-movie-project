package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	var status = fiber.StatusOK
	if data != nil {
		return c.JSON(&dto.DataResponse{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(&dto.DefaultResponse{
		Status:  status,
		Message: message,
	})
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	var status = fiber.StatusNotFound
	return c.Status(status).JSON(&dto.DefaultResponse{
		Status:  status,
		Message: message,
	})
}

// ResponseError : returning json structure for error request
func ResponseError(c *fiber.Ctx, message string, data interface{}) error {
	var status = fiber.StatusInternalServerError
	if data != nil {
		return c.Status(status).JSON(&dto.DataResponse{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

	return c.Status(status).JSON(&dto.DefaultResponse{
		Status:  status,
		Message: message,
	})
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, message string, data interface{}) error {
	var status = fiber.StatusUnauthorized
	if data != nil {
		return c.Status(status).JSON(&dto.DataResponse{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

	return c.Status(status).JSON(&dto.DefaultResponse{
		Status:  status,
		Message: message,
	})
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	var status = fiber.StatusBadRequest
	if data != nil {
		return c.Status(status).JSON(&dto.DataResponse{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

	return c.Status(status).JSON(&dto.DefaultResponse{
		Status:  status,
		Message: message,
	})
}
