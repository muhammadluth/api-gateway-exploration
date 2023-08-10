package handler

import "github.com/gofiber/fiber/v2"

type IPostRouter interface {
	CreatePost(ctx *fiber.Ctx) error
	GetListPost(ctx *fiber.Ctx) error
	GetDetailPost(ctx *fiber.Ctx) error
	UpdatePost(ctx *fiber.Ctx) error
	DeletePost(ctx *fiber.Ctx) error
}
