package src

import (
	"post-backend/model"
)

type IPostMapper interface {
	ToInsertPost(agent, userId string, request model.RequestCreatePost) (post model.Post)
	ToResponseSuccessGetListPost(traceId string, message model.Message, count int, posts []model.Post) (response model.Response)
	ToResponseSuccessGetDetailPost(traceId string, post model.Post) (response model.Response)
	ToUpdatePost(agent string, request model.RequestUpdatePost) (post model.Post)
}
