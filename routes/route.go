package routes

import (
	"antrian/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	queue := r.Group("/api")

	//queue.Get("/", controllers.Index)
	queue.Get("/:id", controllers.Show)
	//queue.Post("/", controllers.Create)
	queue.Put("/:id", controllers.Update)
	queue.Put("/reset/:id", controllers.Reset)
	//queue.Delete("/:id", controllers.Delete)
}