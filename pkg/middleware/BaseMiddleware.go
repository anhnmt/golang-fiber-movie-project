package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
)

func BaseMiddleware(a *fiber.App) {
	serverConfig := config.GetServer()

	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Add Recover
	a.Use(recover.New())

	// Add Icon
	a.Use(favicon.New())

	// Add CORS to each route.
	a.Use(cors.New())

	// Add simple logger.
	if serverConfig.Logger {
		a.Use(logger.New())
	}

	// Add caching.
	// cache.New(),
}
