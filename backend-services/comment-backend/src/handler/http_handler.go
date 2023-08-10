package handler

import (
	"comment-backend/src"

	"github.com/gofiber/fiber/v2"
)

type CommentHttpHandler struct {
	fiberRouter    fiber.Router
	iCommentRouter ICommentRouter
}

func NewCommentHttpHandler(fiberRouter fiber.Router, iCommentRouter ICommentRouter) src.ICommentHttpHandler {
	return &CommentHttpHandler{fiberRouter, iCommentRouter}
}

func (h *CommentHttpHandler) Routers() {
	fiberAppComment := h.fiberRouter.Group("/comment")
	fiberAppComment.Post("/", h.iCommentRouter.CreateComment)
	fiberAppComment.Get("/", h.iCommentRouter.GetListComment)
	fiberAppComment.Get("/:id", h.iCommentRouter.GetDetailComment)
	fiberAppComment.Put("/:id", h.iCommentRouter.UpdateComment)
	fiberAppComment.Delete("/:id", h.iCommentRouter.DeleteComment)
}
