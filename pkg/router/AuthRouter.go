package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/controller"
)

func authRouter(a fiber.Router) {
	oauth := a.Group("/oauth")

	oauth.Post("/token", controller.AuthToken)
}
