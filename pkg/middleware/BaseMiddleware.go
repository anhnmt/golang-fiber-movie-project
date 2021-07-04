package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
)

func BaseMiddleware(a *fiber.App) {
	serverConfig := config.GetServer()

	// Add CORS to each route.
	a.Use(cors.New())

	// Add simple logger.
	if serverConfig.Logger {
		a.Use(logger.New())
	}

	// Add caching.
	//cache.New(),
}
