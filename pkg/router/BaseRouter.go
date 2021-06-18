package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func BaseRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":  "Welcome to Fiber Go API!",
			"docs": "/swagger/index.html",
		})
	})

	swaggerRoute(a)

	api := a.Group("/api")
	privateRoute(api)

	// Auth Router
	authRouter(api)

	// 404 Not Found Router
	notFoundRoute(a)
}

func swaggerRoute(a *fiber.App) {
	// Create route group.
	route := a.Group("/swagger")
	route.Get("*", swagger.Handler)
}

func notFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
		},
	)
}
