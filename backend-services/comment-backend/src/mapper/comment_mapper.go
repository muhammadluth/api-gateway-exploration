package mapper

import (
	"comment-backend/app/utils"
	"comment-backend/model"
	"comment-backend/src"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CommentMapper struct {
}

func NewCommentMapper() src.ICommentMapper {
	return &CommentMapper{}
}

func (m *CommentMapper) ToInsertComment(agent, userId string, request model.RequestCreateComment) (comment model.Comment) {
	comment = model.Comment{
		ID:     uuid.NewString(),
		UserID: userId,
		Body:   request.Body,
		Agent:  agent,
		PostID: request.PostID,
	}
	return comment
}

func (m *CommentMapper) ToResponseSuccessGetListComment(traceId string, message model.Message, count int,
	comments []model.Comment) (response model.Response) {
	resListComment := []model.ResponseGetListComment{}
	for _, item := range comments {
		var createdAt, updatedAt *string
		if !item.CreatedAt.IsZero() {
			format := item.CreatedAt.Format(time.RFC3339)
			createdAt = &format
		}
		if !item.UpdatedAt.IsZero() {
			format := item.UpdatedAt.Format(time.RFC3339)
			updatedAt = &format
		}
		data := model.ResponseGetListComment{
			ID:        item.ID,
			Body:      item.Body,
			UserID:    item.UserID,
			Agent:     item.Agent,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		resListComment = append(resListComment, data)
	}
	return model.ResponseSuccessPaginationData(message, http.StatusOK, count, resListComment)
}

func (m *CommentMapper) ToResponseSuccessGetDetailComment(traceId string, comment model.Comment) (response model.Response) {
	// resListComment := []model.ResponseGetListComment{}
	// for _, item := range post.Comments {
	// 	var createdAt, updatedAt *string
	// 	if !item.CreatedAt.IsZero() {
	// 		format := item.CreatedAt.Format(time.RFC3339)
	// 		createdAt = &format
	// 	}
	// 	if !item.UpdatedAt.IsZero() {
	// 		format := item.UpdatedAt.Format(time.RFC3339)
	// 		updatedAt = &format
	// 	}
	// 	data := model.ResponseGetListComment{
	// 		ID:        item.ID,
	// 		Body:      item.Body,
	// 		UserID:    item.UserID,
	// 		Agent:     item.Agent,
	// 		CreatedAt: createdAt,
	// 		UpdatedAt: updatedAt,
	// 	}
	// 	resListComment = append(resListComment, data)
	// }

	var createdAt, updatedAt *string
	if !comment.CreatedAt.IsZero() {
		format := comment.CreatedAt.Format(time.RFC3339)
		createdAt = &format
	}
	if !comment.UpdatedAt.IsZero() {
		format := comment.UpdatedAt.Format(time.RFC3339)
		updatedAt = &format
	}
	resDetailComment := model.ResponseGetDetailComment{
		ID:        comment.ID,
		Body:      comment.Body,
		UserID:    comment.UserID,
		Agent:     comment.Agent,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if comment.Post != nil {
		var createdAt, updatedAt *string
		if !comment.Post.CreatedAt.IsZero() {
			format := comment.Post.CreatedAt.Format(time.RFC3339)
			createdAt = &format
		}
		if !comment.Post.UpdatedAt.IsZero() {
			format := comment.Post.UpdatedAt.Format(time.RFC3339)
			updatedAt = &format
		}
		resDetailPost := model.ResponseGetDetailPost{
			ID:        comment.Post.ID,
			Title:     comment.Post.Title,
			Body:      comment.Post.Body,
			UserID:    comment.Post.UserID,
			Agent:     comment.Post.Agent,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		resDetailComment.Post = resDetailPost
	}

	return model.ResponseSuccessData(http.StatusOK, resDetailComment)
}

func (m *CommentMapper) ToUpdateComment(agent string, request model.RequestUpdateComment) (comment model.Comment) {
	comment = model.Comment{
		Agent:     agent,
		UpdatedAt: time.Now(),
	}
	if utils.IsFulfilled(request.Body) {
		comment.Body = request.Body
	}
	return comment
}
