package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(ctx *fiber.Ctx) error {
	cors := cors.New(cors.Config{
		AllowOrigins: "*",
	})

	return cors(ctx)
}