package router

import "github.com/gofiber/fiber/v2"

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":  "Welcome to Fiber Go API!",
			"docs": "/swagger/index.html",
		})
	})

	swaggerRoute(a)
	publicRoute(a)
	privateRoute(a)
	notFoundRoute(a)
}

func notFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}
