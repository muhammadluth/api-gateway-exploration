package server

import (
	"comment-backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type SetupServer struct {
	fiberApp      *fiber.App
	svcProperties model.ServiceProperties
}

func NewSetupServer(svcProperties model.ServiceProperties) SetupServer {
	fiberApp := fiber.New(
		fiber.Config{
			CaseSensitive: true,
			Concurrency:   svcProperties.ServicePoolSizeConnection,
		},
	)
	return SetupServer{fiberApp, svcProperties}
}

func (s *SetupServer) InitServerConfiguration() *fiber.App {
	s.fiberApp.Use(etag.New())
	s.fiberApp.Use(compress.New())
	s.fiberApp.Use(requestid.New())
	s.fiberApp.Use(recover.New())
	s.fiberApp.Use(cors.New(cors.Config{
		Next:         cors.ConfigDefault.Next,
		AllowOrigins: "*",
		AllowMethods: fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions),
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "*",
	}))

	// HEALTH CHECK
	s.fiberApp.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My API!"})
	})

	return s.fiberApp
}

func (s *SetupServer) InitServer() {
	svcPort := s.svcProperties.ServicePort
	fmt.Printf("Listening on port : %d\n", svcPort)
	fmt.Printf("Ready to serve\n")
	s.fiberApp.Listen(fmt.Sprintf(":%d", svcPort))
}
