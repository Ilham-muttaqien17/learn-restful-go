package routes

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {

	// Api routes
	apiRoutes := app.Group("/api")

	RegisterBookRoutes(apiRoutes)
}