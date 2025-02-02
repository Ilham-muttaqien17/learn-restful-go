package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compress(ctx *fiber.Ctx) error {
	compress := compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})

	return compress(ctx)
}