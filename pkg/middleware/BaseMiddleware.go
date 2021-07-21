package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
)

func BaseMiddleware(a *fiber.App) {
	serverConfig := config.GetServer()

	a.Use(compress.New(compress.Config{
		Next:  nil,
		Level: compress.LevelBestSpeed, // 1
	}))

	// Add CORS to each route.
	a.Use(cors.New())

	// Add simple logger.
	if serverConfig.Logger {
		a.Use(logger.New())
	}

	// Add caching.
	//cache.New(),
}
