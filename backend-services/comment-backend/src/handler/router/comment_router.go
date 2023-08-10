package router

import (
	"comment-backend/app/log"
	"comment-backend/app/utils"
	"comment-backend/model"
	"comment-backend/model/constant"
	"comment-backend/src"
	"comment-backend/src/handler"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CommentRouter struct {
	iCommentUsecase src.ICommentUsecase
}

func NewCommentRouter(iCommentUsecase src.ICommentUsecase) handler.ICommentRouter {
	return &CommentRouter{iCommentUsecase}
}

func (r *CommentRouter) CreateComment(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, nil, nil, ctx.Request().Body())
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iCommentUsecase.CreateComment(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *CommentRouter) GetListComment(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	query := new(model.QueryGetListComment)
	if err := ctx.QueryParser(query); err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	queryAsByteJSON, _ := json.Marshal(query)
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, nil, queryAsByteJSON, nil)
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iCommentUsecase.GetListComment(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *CommentRouter) GetDetailComment(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamComment)
	if err := ctx.ParamsParser(param); err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	paramAsByteJSON, _ := json.Marshal(param)
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, paramAsByteJSON, nil, nil)
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iCommentUsecase.GetDetailComment(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *CommentRouter) UpdateComment(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamComment)
	if err := ctx.ParamsParser(param); err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	paramAsByteJSON, _ := json.Marshal(param)
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, paramAsByteJSON, nil, ctx.Request().Body())
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iCommentUsecase.UpdateComment(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *CommentRouter) DeleteComment(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamComment)
	if err := ctx.ParamsParser(param); err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	paramAsByteJSON, _ := json.Marshal(param)
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, paramAsByteJSON, nil, nil)
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iCommentUsecase.DeleteComment(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}
