package router

import (
	"github.com/gofiber/fiber/v2"
	"go_blog/controller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/blogs", controller.BlogList)
	api.Post("/blogs", controller.BlogCreate)
	api.Put("/blogs/:id", controller.BlogUpdate)
	api.Delete("/blogs/:id", controller.BlogDelete)
}
