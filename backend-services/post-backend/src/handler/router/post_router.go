package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"post-backend/app/log"
	"post-backend/app/utils"
	"post-backend/model"
	"post-backend/model/constant"
	"post-backend/src"
	"post-backend/src/handler"

	"github.com/gofiber/fiber/v2"
)

type PostRouter struct {
	iPostUsecase src.IPostUsecase
}

func NewPostRouter(iPostUsecase src.IPostUsecase) handler.IPostRouter {
	return &PostRouter{iPostUsecase}
}

func (r *PostRouter) CreatePost(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	message, err := utils.GenerateMessage(traceId, ctx, http.Header{}, nil, nil, ctx.Request().Body())
	if err != nil {
		log.Error(err, traceId)
		response := model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iPostUsecase.CreatePost(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *PostRouter) GetListPost(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	query := new(model.QueryGetListPost)
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
	response := r.iPostUsecase.GetListPost(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *PostRouter) GetDetailPost(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamPost)
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
	response := r.iPostUsecase.GetDetailPost(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *PostRouter) UpdatePost(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamPost)
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
	response := r.iPostUsecase.UpdatePost(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *PostRouter) DeletePost(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(constant.LOCALS_TRACE_ID))
	param := new(model.ParamPost)
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
	response := r.iPostUsecase.DeletePost(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}
