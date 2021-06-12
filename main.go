package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"github.com/xdorro/golang-fiber-base-project/pkg/middleware"
	"github.com/xdorro/golang-fiber-base-project/pkg/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := fiber.New()

	cfg := config.GetConfig()

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	router.GeneralRoute(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Printf("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Println(serverAddr)

	if err := app.Listen(serverAddr); err != nil {
		log.Printf("Oops... server is not running! error: %v", err)
	}
}
