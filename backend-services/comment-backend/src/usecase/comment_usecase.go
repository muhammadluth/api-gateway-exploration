package usecase

import (
	"comment-backend/app/log"
	"comment-backend/model"
	"comment-backend/model/constant"
	"comment-backend/src"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CommentUsecase struct {
	svcProperties  model.ServiceProperties
	iCommentMapper src.ICommentMapper
	iCommentRepo   src.ICommentRepo
}

func NewCommentUsecase(svcProperties model.ServiceProperties,
	iCommentMapper src.ICommentMapper, iCommentRepo src.ICommentRepo) src.ICommentUsecase {
	return &CommentUsecase{svcProperties, iCommentMapper, iCommentRepo}
}

func (u *CommentUsecase) CreateComment(traceId string, message model.Message) model.Response {
	request := model.RequestCreateComment{}
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	commentData := u.iCommentMapper.ToInsertComment(message.Header.URL, u.svcProperties.CustomUserID, request)

	if err := u.iCommentRepo.InsertComment(traceId, commentData); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusCreated, "Successfully Create Comment")
}

func (u *CommentUsecase) GetListComment(traceId string, message model.Message) model.Response {
	query := new(model.QueryGetListComment)
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

	if query.UserID == "" {
		query.UserID = u.svcProperties.CustomUserID
	}

	count, comments, err := u.iCommentRepo.GetListCommentByWhereQuery(traceId, queryBeginDate, queryUntilDate, *query)
	if err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	response := u.iCommentMapper.ToResponseSuccessGetListComment(traceId, message, count, comments)

	return response
}

func (u *CommentUsecase) GetDetailComment(traceId string, message model.Message) model.Response {
	param := new(model.ParamComment)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	comment, err := u.iCommentRepo.GetCommentWithPostByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	response := u.iCommentMapper.ToResponseSuccessGetDetailComment(traceId, comment)

	return response
}

func (u *CommentUsecase) UpdateComment(traceId string, message model.Message) model.Response {
	param := new(model.ParamComment)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	request := model.RequestUpdateComment{}
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	commentData, err := u.iCommentRepo.GetCommentWithoutPostByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	} else if commentData.ID == "" {
		log.Error(fmt.Errorf("comment data '%v' is not found", param.ID), traceId)
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	}

	comment := u.iCommentMapper.ToUpdateComment(message.Header.URL, request)

	if err := u.iCommentRepo.UpdateCommentByIDAndUserID(traceId, commentData.ID, commentData.UserID, comment); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusOK, "Successfully Update Comment")
}

func (u *CommentUsecase) DeleteComment(traceId string, message model.Message) model.Response {
	param := new(model.ParamComment)
	if err := json.Unmarshal(message.Param, &param); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusBadRequest, "Invalid Request")
	}

	commentData, err := u.iCommentRepo.GetCommentWithoutPostByIDAndUserID(traceId, param.ID, u.svcProperties.CustomUserID)
	if err == sql.ErrNoRows {
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	} else if err != nil {
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	} else if commentData.ID == "" {
		log.Error(fmt.Errorf("comment data '%v' is not found", param.ID), traceId)
		return model.ResponseErrorDefault(http.StatusNotFound, "Data Not Found")
	}

	if err := u.iCommentRepo.DeleteCommentByIDAndUserID(traceId, commentData.ID, commentData.UserID); err != nil {
		log.Error(err, traceId)
		return model.ResponseErrorDefault(http.StatusInternalServerError, "Server Error")
	}

	return model.ResponseSuccessDefault(http.StatusOK, "Successfully Delete Comment")
}
