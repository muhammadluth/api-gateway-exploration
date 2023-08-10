package src

import (
	"post-backend/model"
	"time"
)

type IPostRepo interface {
	InsertPost(traceId string, data model.Post) (err error)
	GetListPostByWhereQuery(traceId string, whereQueryBeginDate, whereQueryUntilDate time.Time,
		whereQuery model.QueryGetListPost) (count int, data []model.Post, err error)
	GetPostWithCommentsByIDAndUserID(traceId, id, userId string) (data model.Post, err error)
	GetPostWithoutCommentsByIDAndUserID(traceId, id, userId string) (data model.Post, err error)
	UpdatePostByIDAndUserID(traceId, id, userId string, data model.Post) (err error)
	DeletePostByIDAndUserID(traceId, id, userId string) (err error)
}
