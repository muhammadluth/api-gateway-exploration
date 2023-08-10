package model

type RequestCreateComment struct {
	PostID string `json:"post_id"`
	Body   string `json:"body"`
}

type RequestUpdateComment struct {
	Body string `json:"body"`
}

type QueryGetListComment struct {
	PageSize  int    `json:"page_size" query:"page-size"`
	Page      int    `json:"page" query:"page"`
	BeginDate string `json:"begin_date" query:"begin-date"`
	UntilDate string `json:"until_date" query:"until-date"`
	PostID    string `json:"post_id" query:"post-id"`
	Comment   string `json:"comment" query:"comment"`
	UserID    string `json:"user_id" query:"user-id"`
}

type ResponseGetListComment struct {
	ID        string  `json:"id"`
	Body      string  `json:"body"`
	UserID    string  `json:"user_id"`
	Agent     string  `json:"agent"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type (
	ResponseGetDetailComment struct {
		ID        string                `json:"id"`
		Body      string                `json:"body"`
		UserID    string                `json:"user_id"`
		Agent     string                `json:"agent"`
		CreatedAt *string               `json:"created_at"`
		UpdatedAt *string               `json:"updated_at"`
		Post      ResponseGetDetailPost `json:"post"`
	}
	ResponseGetDetailPost struct {
		ID        string  `json:"id"`
		Title     string  `json:"title"`
		Body      string  `json:"body"`
		UserID    string  `json:"user_id"`
		Agent     string  `json:"agent"`
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}
)

type ParamComment struct {
	ID string `json:"id" params:"id"`
}
