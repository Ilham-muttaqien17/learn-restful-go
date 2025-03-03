package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestId(ctx *fiber.Ctx) error {
	requestId := requestid.New(requestid.Config{
		ContextKey: "requestId",
		Header:     "X-Request-Id",
		Generator:  utils.UUIDv4,
	})

	return requestId(ctx)
}
