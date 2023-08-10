package mapper

import (
	"net/http"
	"post-backend/app/utils"
	"post-backend/model"
	"post-backend/src"
	"time"

	"github.com/google/uuid"
)

type PostMapper struct {
}

func NewPostMapper() src.IPostMapper {
	return &PostMapper{}
}

func (m *PostMapper) ToInsertPost(agent, userId string, request model.RequestCreatePost) (post model.Post) {
	post = model.Post{
		ID:     uuid.NewString(),
		UserID: userId,
		Title:  request.Title,
		Body:   request.Body,
		Agent:  agent,
	}
	return post
}

func (m *PostMapper) ToResponseSuccessGetListPost(traceId string, message model.Message, count int,
	posts []model.Post) (response model.Response) {
	resListPost := []model.ResponseGetListPost{}
	for _, item := range posts {
		var createdAt, updatedAt *string
		if !item.CreatedAt.IsZero() {
			format := item.CreatedAt.Format(time.RFC3339)
			createdAt = &format
		}
		if !item.UpdatedAt.IsZero() {
			format := item.UpdatedAt.Format(time.RFC3339)
			updatedAt = &format
		}
		data := model.ResponseGetListPost{
			ID:        item.ID,
			Title:     item.Title,
			Body:      item.Body,
			UserID:    item.UserID,
			Agent:     item.Agent,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		resListPost = append(resListPost, data)
	}
	return model.ResponseSuccessPaginationData(message, http.StatusOK, count, resListPost)
}

func (m *PostMapper) ToResponseSuccessGetDetailPost(traceId string, post model.Post) (response model.Response) {
	resListComment := []model.ResponseGetListComment{}
	for _, item := range post.Comments {
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

	var createdAt, updatedAt *string
	if !post.CreatedAt.IsZero() {
		format := post.CreatedAt.Format(time.RFC3339)
		createdAt = &format
	}
	if !post.UpdatedAt.IsZero() {
		format := post.UpdatedAt.Format(time.RFC3339)
		updatedAt = &format
	}
	resDetailPost := model.ResponseGetDetailPost{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		UserID:    post.UserID,
		Agent:     post.Agent,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Comments:  resListComment,
	}
	return model.ResponseSuccessData(http.StatusOK, resDetailPost)
}

func (m *PostMapper) ToUpdatePost(agent string, request model.RequestUpdatePost) (post model.Post) {
	post = model.Post{
		Agent:     agent,
		UpdatedAt: time.Now(),
	}
	if utils.IsFulfilled(request.Title) {
		post.Title = request.Title
	}
	if utils.IsFulfilled(request.Body) {
		post.Body = request.Body
	}
	return post
}
