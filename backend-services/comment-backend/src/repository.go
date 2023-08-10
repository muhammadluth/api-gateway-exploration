package src

import (
	"comment-backend/model"
	"time"
)

type ICommentRepo interface {
	InsertComment(traceId string, data model.Comment) (err error)
	GetListCommentByWhereQuery(traceId string, whereQueryBeginDate, whereQueryUntilDate time.Time,
		whereQuery model.QueryGetListComment) (count int, data []model.Comment, err error)
	GetCommentWithPostByIDAndUserID(traceId, id, userId string) (data model.Comment, err error)
	GetCommentWithoutPostByIDAndUserID(traceId, id, userId string) (data model.Comment, err error)
	UpdateCommentByIDAndUserID(traceId, id, userId string, data model.Comment) (err error)
	DeleteCommentByIDAndUserID(traceId, id, userId string) (err error)
}
