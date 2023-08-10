package repository

import (
	"comment-backend/app/log"
	"comment-backend/model"
	"comment-backend/src"
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type CommentRepo struct {
	db *bun.DB
}

func NewCommentRepo(db *bun.DB) src.ICommentRepo {
	return &CommentRepo{db}
}

func (r *CommentRepo) InsertComment(traceId string, data model.Comment) (err error) {
	res, err := r.db.NewInsert().Model(&data).Exec(context.Background())
	if err != nil {
		log.Error(err, traceId)
	} else if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 || err != nil {
		err = fmt.Errorf("no rows affected : '%v'", rowsAffected)
		log.Error(err, traceId)
	}
	return err
}

func (r *CommentRepo) GetListCommentByWhereQuery(traceId string, whereQueryBeginDate, whereQueryUntilDate time.Time,
	whereQuery model.QueryGetListComment) (count int, data []model.Comment, err error) {
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

	if whereQuery.PostID != "" {
		query.Where("post_id = ?", whereQuery.PostID)
	}

	if whereQuery.Comment != "" {
		query.Where("body ILIKE ?", "%"+whereQuery.Comment+"%")
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

func (r *CommentRepo) GetCommentWithPostByIDAndUserID(traceId, id, userId string) (data model.Comment, err error) {
	err = r.db.NewSelect().Model(&data).Relation("Post").
		Where("p_c.user_id = ?", userId).
		Where("p_c.id = ?", id).
		Scan(context.Background())
	if err != nil {
		log.Error(err, traceId)
	}
	return data, err
}

func (r *CommentRepo) GetCommentWithoutPostByIDAndUserID(traceId, id, userId string) (data model.Comment, err error) {
	err = r.db.NewSelect().Model(&data).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		log.Error(err, traceId)
	}
	return data, err
}

func (r *CommentRepo) UpdateCommentByIDAndUserID(traceId, id, userId string, data model.Comment) (err error) {
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

func (r *CommentRepo) DeleteCommentByIDAndUserID(traceId, id, userId string) (err error) {
	res, err := r.db.NewDelete().Model(&model.Comment{}).
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
