package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	if data != nil {
		return c.JSON(&dto.DataResponse{
			Status:  fiber.StatusOK,
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(&dto.DefaultResponse{
		Status:  fiber.StatusOK,
		Message: message,
	})
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	return c.JSON(&dto.DefaultResponse{
		Status:  fiber.StatusNotFound,
		Message: message,
	})
}

// ResponseError : returning json structure for error request
func ResponseError(c *fiber.Ctx, message string, data interface{}) error {
	if data != nil {
		return c.JSON(&dto.DataResponse{
			Status:  fiber.StatusInternalServerError,
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(&dto.DefaultResponse{
		Status:  fiber.StatusInternalServerError,
		Message: message,
	})
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, message string, data interface{}) error {
	if data != nil {
		return c.JSON(&dto.DataResponse{
			Status:  fiber.StatusUnauthorized,
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(&dto.DataResponse{
		Status:  fiber.StatusUnauthorized,
		Message: message,
	})
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	if data != nil {
		return c.JSON(&dto.DataResponse{
			Status:  fiber.StatusBadRequest,
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(&dto.DataResponse{
		Status:  fiber.StatusBadRequest,
		Message: message,
	})
}
