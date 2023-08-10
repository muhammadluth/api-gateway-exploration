package src

import "comment-backend/model"

type ICommentUsecase interface {
	CreateComment(traceId string, message model.Message) model.Response
	GetListComment(traceId string, message model.Message) model.Response
	GetDetailComment(traceId string, message model.Message) model.Response
	UpdateComment(traceId string, message model.Message) model.Response
	DeleteComment(traceId string, message model.Message) model.Response
}
