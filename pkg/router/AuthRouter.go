package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-movie-project/app/controller"
	"github.com/xdorro/golang-fiber-movie-project/pkg/middleware"
)

func authRouter(a fiber.Router) {
	oauth := a.Group("/oauth")

	authController := controller.NewAuthController()
	oauth.Post("/token", authController.AuthToken)

	protected := oauth.Use(middleware.Protected())
	protected.Get("/restricted", authController.Restricted)
	protected.Get("/current_user", authController.CurrentUser)
	protected.Post("/change_password", authController.ChangePassword)
	protected.Post("/update_profile", authController.UpdateProfile)
}
