package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func swaggerRoute(a *fiber.App) {
	// Create route group.
	route := a.Group("/swagger")
	route.Get("*", swagger.Handler)
}
