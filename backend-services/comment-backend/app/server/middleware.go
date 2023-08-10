package server

import (
	"comment-backend/app/utils"
	"comment-backend/model/constant"

	"github.com/gofiber/fiber/v2"
)

func (s *SetupServer) InitServerMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		traceId := ctx.Get("trace_id")
		if traceId == "" {
			traceId = utils.CreateTraceID()
		}

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		ctx.Locals(constant.LOCALS_TRACE_ID, traceId)

		return ctx.Next()
	}
}
