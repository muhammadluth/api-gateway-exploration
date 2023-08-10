package utils

import (
	"net/http"
	"post-backend/model"
	"post-backend/model/constant"

	"github.com/gofiber/fiber/v2"
)

func GenerateMessage(traceId string, ctx *fiber.Ctx, header http.Header, param, query, rawReqBody []byte) (message model.Message, err error) {
	authData, _ := ctx.Locals(constant.LOCAL_AUTH_DATA).(model.AuthData)
	message = model.Message{
		Header: model.Header{
			URL:    string(ctx.BaseURL() + ctx.OriginalURL()),
			Query:  string(ctx.Context().QueryArgs().QueryString()),
			Header: header,
		},
		Body:     rawReqBody,
		Param:    param,
		Query:    query,
		AuthData: authData,
	}
	return message, err
}
