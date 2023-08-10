package repository

import (
	"context"
	"fmt"
	"post-backend/app/log"
	"post-backend/model"
	"post-backend/src"
	"time"

	"github.com/uptrace/bun"
)

type PostRepo struct {
	db *bun.DB
}

func NewPostRepo(db *bun.DB) src.IPostRepo {
	return &PostRepo{db}
}

func (r *PostRepo) InsertPost(traceId string, data model.Post) (err error) {
	res, err := r.db.NewInsert().Model(&data).Exec(context.Background())
	if err != nil {
		log.Error(err, traceId)
	} else if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 || err != nil {
		err = fmt.Errorf("no rows affected : '%v'", rowsAffected)
		log.Error(err, traceId)
	}
	return err
}

func (r *PostRepo) GetListPostByWhereQuery(traceId string, whereQueryBeginDate, whereQueryUntilDate time.Time,
	whereQuery model.QueryGetListPost) (count int, data []model.Post, err error) {
	query := r.db.NewSelect().Model(&data)

	if !whereQueryBeginDate.IsZero() {
		query.Where("updated_at > ?", whereQueryBeginDate.AddDate(0, 0, -1))
	}

	if !whereQueryUntilDate.IsZero() {
		query.Where("updated_at < ?", whereQueryUntilDate.AddDate(0, 0, 1))
	}

	if whereQuery.UserID != "" {
		query.Where("user_id = ?", whereQuery.UserID)
	}

	if whereQuery.Post != "" {
		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q.Where("title ILIKE ?", "%"+whereQuery.Post+"%")
			q.WhereOr("body ILIKE ?", "%"+whereQuery.Post+"%")
			return q
		})
	}

	if whereQuery.Page > 0 {
		whereQuery.Page = (whereQuery.Page - 1) * whereQuery.PageSize
	}
	count, err = query.Limit(whereQuery.PageSize).Offset(whereQuery.Page).OrderExpr("updated_at DESC").ScanAndCount(context.Background())
	if err != nil {
		log.Error(err, traceId)
	}
	return count, data, err
}

func (r *PostRepo) GetPostWithCommentsByIDAndUserID(traceId, id, userId string) (data model.Post, err error) {
	err = r.db.NewSelect().Model(&data).Relation("Comments").
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		log.Error(err, traceId)
	}
	return data, err
}

func (r *PostRepo) GetPostWithoutCommentsByIDAndUserID(traceId, id, userId string) (data model.Post, err error) {
	err = r.db.NewSelect().Model(&data).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		log.Error(err, traceId)
	}
	return data, err
}

func (r *PostRepo) UpdatePostByIDAndUserID(traceId, id, userId string, data model.Post) (err error) {
	res, err := r.db.NewUpdate().Model(&data).OmitZero().
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		log.Error(err, traceId)
	} else if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 || err != nil {
		err = fmt.Errorf("no rows affected : '%v'", rowsAffected)
		log.Error(err, traceId)
	}
	return err
}

func (r *PostRepo) DeletePostByIDAndUserID(traceId, id, userId string) (err error) {
	res, err := r.db.NewDelete().Model(&model.Post{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		log.Error(err, traceId)
	} else if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 || err != nil {
		err = fmt.Errorf("no rows affected : '%v'", rowsAffected)
		log.Error(err, traceId)
	}
	return err
}
