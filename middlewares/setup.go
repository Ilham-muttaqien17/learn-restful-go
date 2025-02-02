package middlewares

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	// Compress
	app.Use(Compress)
	// CORS
	app.Use(Cors)
	// Helmet
	app.Use(Helmet)
	// Rate Limiter
	app.Use(RateLimiter)
	// Request Id
	app.Use(RequestId)
}