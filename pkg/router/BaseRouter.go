package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func BaseRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":  "Welcome to Fiber Go API!",
			"docs": "/swagger/index.html",
		})
	})

	// Create route group.
	app.Get("/swagger/*", swagger.Handler)

	api := app.Group("/api")
	privateRoute(api)

	// Auth Router
	authRouter(api)

	// 404 Not Found Router
	notFoundRoute(app)
}

func notFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return util.ResponseNotFound("Đường dẫn không tồn tại")
		},
	)
}
