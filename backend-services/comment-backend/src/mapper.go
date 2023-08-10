package src

import (
	"comment-backend/model"
)

type ICommentMapper interface {
	ToInsertComment(agent, userId string, request model.RequestCreateComment) (comment model.Comment)
	ToResponseSuccessGetListComment(traceId string, message model.Message, count int, comments []model.Comment) (response model.Response)
	ToResponseSuccessGetDetailComment(traceId string, comment model.Comment) (response model.Response)
	ToUpdateComment(agent string, request model.RequestUpdateComment) (comment model.Comment)
}
