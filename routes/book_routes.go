package routes

import (
	"github.com/Ilham-muttaqien17/learn-restful-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(router fiber.Router) {
	bookController := controllers.RegisterBookController()
	bookRoutes := router.Group("/books")

	bookRoutes.Get("/", bookController.Index)
	bookRoutes.Get("/:id", bookController.Show)
	bookRoutes.Post("/", bookController.Create)
	bookRoutes.Patch("/:id", bookController.Update)
	bookRoutes.Delete("/:id", bookController.Destroy)
}
