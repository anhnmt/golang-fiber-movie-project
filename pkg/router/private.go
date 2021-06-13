package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
)

func privateRoute(a *fiber.App) {
	api := a.Group("/api")

	// Tags Controller
	tags := api.Group("/tags")
	tags.Get("/", controller.FindAllTags)
	tags.Post("/", controller.CreateNewTag)
	tags.Get("/:id", controller.FindTagById)
	tags.Put("/:id", controller.UpdateTagById)
	tags.Delete("/:id", controller.DeleteTagById)
}
