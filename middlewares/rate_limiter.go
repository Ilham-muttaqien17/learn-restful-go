package middlewares

import (
	"time"

	"github.com/Ilham-muttaqien17/learn-restful-go/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(ctx *fiber.Ctx) error {
	rateLimiter := limiter.New(limiter.Config{
		Max:        60,
		Expiration: time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return "limiter-" + c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(map[string]interface{}{
				"message": "Too many request, try again later!",
			})
		},
		Storage: config.RedisStore,
	})

	return rateLimiter(ctx)
}
