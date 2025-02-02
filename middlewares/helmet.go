package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Helmet(ctx *fiber.Ctx) error {
	helmet := helmet.New()

	return helmet(ctx)
}