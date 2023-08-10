package handler

import (
	"post-backend/src"

	"github.com/gofiber/fiber/v2"
)

type PostHttpHandler struct {
	fiberRouter fiber.Router
	iPostRouter IPostRouter
}

func NewPostHttpHandler(fiberRouter fiber.Router, iPostRouter IPostRouter) src.IPostHttpHandler {
	return &PostHttpHandler{fiberRouter, iPostRouter}
}

func (h *PostHttpHandler) Routers() {
	fiberAppPost := h.fiberRouter.Group("/post")
	fiberAppPost.Post("/", h.iPostRouter.CreatePost)
	fiberAppPost.Get("/", h.iPostRouter.GetListPost)
	fiberAppPost.Get("/:id", h.iPostRouter.GetDetailPost)
	fiberAppPost.Put("/:id", h.iPostRouter.UpdatePost)
	fiberAppPost.Delete("/:id", h.iPostRouter.DeletePost)
}
