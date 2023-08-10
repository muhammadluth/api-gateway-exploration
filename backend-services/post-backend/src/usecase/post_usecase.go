package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"post-backend/app/log"
	"post-backend/model"
	"post-backend/model/constant"
	"post-backend/src"
	"time"
)

type PostUsecase struct {
	svcProperties model.ServiceProperties
	iPostMapper   src.IPostMapper
	iPostRepo     src.IPostRepo
}

func NewPostUsecase(svcProperties model.ServiceProperties,
	iPostMapper src.IPostMapper, iPostRepo src.IPostRepo) src.IPostUsecase {
	return &PostUsecase{svcProperties, iPostMapper, iPostRepo}
}

func (u *PostUsecase) CreatePost(traceId string, message model.Message) model.Response {
	request := model.RequestCreatePost{}
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	postData := u.iPostMapper.ToInsertPost(message.Header.URL, u.svcProperties.CustomUserID, request)

	if err := u.iPostRepo.InsertPost(traceId, postData); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusCreated, "Successfully Create Post")
}

func (u *PostUsecase) GetListPost(traceId string, message model.Message) model.Response {
	query := new(model.QueryGetListPost)
	if err := json.Unmarshal(message.Query, &query); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	queryBeginDate := time.Time{}
	if query.BeginDate != "" {
		queryBeginDate, _ = time.ParseInLocation(constant.YYYY_MM_DD, query.BeginDate, u.svcProperties.ServiceTimezone)
	}

	queryUntilDate := time.Time{}
	if query.UntilDate != "" {
		queryUntilDate, _ = time.ParseInLocation(constant.YYYY_MM_DD, query.UntilDate, u.svcProperties.ServiceTimezone)
	}

	count, posts, err := u.iPostRepo.GetListPostByWhereQuery(traceId, queryBeginDate, queryUntilDate, *query)
	if err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	response := u.iPostMapper.ToResponseSuccessGetListPost(traceId, message, count, posts)

	return response
}

func (u *PostUsecase) GetDetailPost(traceId string, message model.Message) model.Response {
	param := new(model.ParamPost)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	post, err := u.iPostRepo.GetPostWithCommentsByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	response := u.iPostMapper.ToResponseSuccessGetDetailPost(traceId, post)

	return response
}

func (u *PostUsecase) UpdatePost(traceId string, message model.Message) model.Response {
	param := new(model.ParamPost)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	request := model.RequestUpdatePost{}
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	postData, err := u.iPostRepo.GetPostWithoutCommentsByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	} else if postData.ID == "" {
		log.Error(fmt.Errorf("post data '%v' is not found", param.ID), traceId)
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	}

	post := u.iPostMapper.ToUpdatePost(message.Header.URL, request)

	if err := u.iPostRepo.UpdatePostByIDAndUserID(traceId, postData.ID, postData.UserID, post); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusOK, "Successfully Update Post")
}

func (u *PostUsecase) DeletePost(traceId string, message model.Message) model.Response {
	param := new(model.ParamPost)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	postData, err := u.iPostRepo.GetPostWithoutCommentsByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	} else if postData.ID == "" {
		log.Error(fmt.Errorf("post data '%v' is not found", param.ID), traceId)
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	}

	if err := u.iPostRepo.DeletePostByIDAndUserID(traceId, postData.ID, postData.UserID); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusOK, "Successfully Delete Post")
}
