package src

import "post-backend/model"

type IPostUsecase interface {
	CreatePost(traceId string, message model.Message) model.Response
	GetListPost(traceId string, message model.Message) model.Response
	GetDetailPost(traceId string, message model.Message) model.Response
	UpdatePost(traceId string, message model.Message) model.Response
	DeletePost(traceId string, message model.Message) model.Response
}
