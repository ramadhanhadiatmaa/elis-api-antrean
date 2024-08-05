package routes

import (
	"antrian/controllers"
	"antrian/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	queue := r.Group("/api")

	queue.Get("/", middlewares.AuthMiddleware, controllers.Index)
	queue.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	queue.Post("/", middlewares.AuthMiddleware, controllers.Create)
	queue.Put("/up/:id", middlewares.AuthMiddleware, controllers.Update)
	queue.Put("/re/:id", middlewares.AuthMiddleware, controllers.Reset)
	queue.Delete("/:id", middlewares.AuthMiddleware, controllers.Delete)
}