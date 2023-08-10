package handler

import "github.com/gofiber/fiber/v2"

type ICommentRouter interface {
	CreateComment(ctx *fiber.Ctx) error
	GetListComment(ctx *fiber.Ctx) error
	GetDetailComment(ctx *fiber.Ctx) error
	UpdateComment(ctx *fiber.Ctx) error
	DeleteComment(ctx *fiber.Ctx) error
}
