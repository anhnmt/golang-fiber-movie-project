package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
)

// ResponseSuccess : returning json structur for success request
func ResponseSuccess(c *fiber.Ctx, message string) error {
	return c.JSON(&dto.DefaultResponse{
		Status:  200,
		Message: message,
	})
}

// ResponseSuccessData : returning json structur for success request
func ResponseSuccessData(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(&dto.DataResponse{
		Status:  200,
		Message: message,
		Data:    data,
	})
}

// ResponseNotFound : returning json structur for notfound request
func ResponseNotFound(c *fiber.Ctx, message string) error {
	return c.JSON(&dto.DefaultResponse{
		Status:  404,
		Message: message,
	})
}

// ResponseErrorData : returning json structur for error request
func ResponseErrorData(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(&dto.DataResponse{
		Status:  500,
		Message: message,
		Data:    data,
	})
}

// ResponseError : returning json structur for error request
func ResponseError(c *fiber.Ctx, message string) error {
	return c.JSON(&dto.DefaultResponse{
		Status:  500,
		Message: message,
	})
}

// ResponseUnauthenticated : returning json structur for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(&dto.DataResponse{
		Status:  403,
		Message: message,
		Data:    data,
	})
}

// ResponseValidationError : returning json structur for validation error request
func ResponseValidationError(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(&dto.DataResponse{
		Status:  304,
		Message: message,
		Data:    data,
	})
}
