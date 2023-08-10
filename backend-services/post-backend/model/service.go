package model

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	ServiceProperties struct {
		ServiceName               string         `json:"service_name" validate:"required"`
		ServicePort               int            `json:"service_port" validate:"numeric,required"`
		ServicePoolSizeConnection int            `json:"service_pool_size_connection" validate:"numeric,required"`
		ServiceTimezone           *time.Location `json:"service_timezone" validate:"required"`
		Database                  DBConfig       `json:"database" validate:"required,dive"`
		CustomUserID              string         `json:"custom_user_id"`
	}

	DBConfig struct {
		IP       string `json:"host" validate:"required"`
		Port     string `json:"port" validate:"numeric,required"`
		User     string `json:"user" validate:"required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
	}
)

type (
	HttpSenderMethod struct {
		URL      string                   `json:"url"`
		Method   string                   `json:"method"`
		Timeout  time.Duration            `json:"timeout"`
		Request  RequestHttpSenderMethod  `json:"request"`
		Response ResponseHttpSenderMethod `json:"response"`
	}

	RequestHttpSenderMethod struct {
		RequestHeader http.Header `json:"request_header"`
		RequestBody   []byte      `json:"request_body"`
	}

	ResponseHttpSenderMethod struct {
		ResponseHeader http.Header `json:"response_header"`
		ResponseBody   []byte      `json:"response_body"`
	}
)

type (
	Message struct {
		Header   Header   `json:"header"`
		Body     []byte   `json:"body"`
		Param    []byte   `json:"param"`
		Query    []byte   `json:"query"`
		AuthData AuthData `json:"auth_data,omitempty"`
	}
	Header struct {
		URL    string      `json:"url"`
		Query  string      `json:"query"`
		Header http.Header `json:"header"`
	}
	AuthData struct {
		UserID string   `json:"user_id"`
		Name   string   `json:"name"`
		Roles  []string `json:"roles"`
	}
)

type Response struct {
	Status int         `json:"status"`
	Header http.Header `json:"header"`
	Body   interface{} `json:"body"`
}

type (
	ResponseDefault struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	ResponseSuccessPagination struct {
		Count       int         `json:"count"`
		Next        *string     `json:"next"`
		Previous    *string     `json:"previous"`
		CurrentPage int         `json:"current_page"`
		TotalPage   int         `json:"total_page"`
		Results     interface{} `json:"results"`
	}

	ResponseSuccessList struct {
		Count   int         `json:"count"`
		Results interface{} `json:"results"`
	}

	ResponseSuccessSingle struct {
		Result interface{} `json:"result"`
	}
)

func ResponseErrorDefault(status int, message string) Response {
	resBody := ResponseDefault{
		Status:  status,
		Message: message,
	}
	return Response{status, http.Header{}, resBody}
}

func ResponseSuccessDefault(status int, message string) Response {
	resBody := ResponseDefault{
		Status:  status,
		Message: message,
	}
	return Response{status, http.Header{}, resBody}
}

func ResponseSuccessCountAndData(status, count int, data interface{}) Response {
	resBody := ResponseSuccessList{
		Count:   count,
		Results: data,
	}
	return Response{status, http.Header{}, resBody}
}

func ResponseSuccessData(status int, data interface{}) Response {
	resBody := ResponseSuccessSingle{
		Result: data,
	}
	return Response{status, http.Header{}, resBody}
}

func ResponseSuccessPaginationData(message Message, status int, count int, data interface{}) Response {
	var next, previous *string
	var pageSize, page float64

	sQueryParams := strings.Split(message.Header.Query, "&")
	for _, item := range sQueryParams {
		sKeyValue := strings.Split(item, "=")
		switch sKeyValue[0] {
		case "page-size", "page_size":
			pageSize, _ = strconv.ParseFloat(sKeyValue[1], 64)
		case "page":
			page, _ = strconv.ParseFloat(sKeyValue[1], 64)
		}
	}

	var totalPage float64
	if pageSize > 0 && page > 0 {
		if pageSize > 0 {
			totalPage = math.Ceil(float64(count) / pageSize)
			if totalPage <= 1 {
				totalPage = 1
			}
		}

		if (page+1) > 0 && (page+1) <= totalPage {
			nextPage := strings.Replace(message.Header.URL, fmt.Sprintf("page=%v", page), fmt.Sprintf("page=%v", page+1), 1)
			next = &nextPage
		}

		if (page - 1) > 0 {
			previousPage := strings.Replace(message.Header.URL, fmt.Sprintf("page=%v", page), fmt.Sprintf("page=%v", page-1), 1)
			previous = &previousPage
		}
	}
	if totalPage == 0 {
		totalPage = 1
	}
	if int(page) == 0 {
		page = 1
	}
	resBody := ResponseSuccessPagination{
		CurrentPage: int(page),
		TotalPage:   int(totalPage),
		Count:       count,
		Next:        next,
		Previous:    previous,
		Results:     data,
	}
	return Response{status, http.Header{}, resBody}
}
