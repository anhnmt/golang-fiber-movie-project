package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(message string, data interface{}) error {
	status := fiber.StatusOK

	return jsonResponse(status, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(message string) error {
	status := fiber.StatusNotFound

	return jsonResponse(status, message, nil)
}

// ResponseError : returning json structure for error request
func ResponseError(message string, data interface{}) error {
	status := fiber.StatusInternalServerError

	return jsonResponse(status, message, data)
}

// ResponseUnauthenticated : returning json structure for validation error request
func ResponseUnauthenticated(message string, data interface{}) error {
	status := fiber.StatusUnauthorized

	return jsonResponse(status, message, data)
}

// ResponseBadRequest : returning json structure for validation error request
func ResponseBadRequest(message string, data interface{}) error {
	status := fiber.StatusBadRequest

	return jsonResponse(status, message, data)
}

func jsonResponse(status int, message string, data interface{}) error {
	msg, _ := JSONMarshal(&response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})

	return fiber.NewError(status, string(msg))
}
