package model

type RequestCreatePost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type RequestUpdatePost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type QueryGetListPost struct {
	PageSize  int    `json:"page_size" query:"page-size"`
	Page      int    `json:"page" query:"page"`
	BeginDate string `json:"begin_date" query:"begin-date"`
	UntilDate string `json:"until_date" query:"until-date"`
	Post      string `json:"post" query:"post"`
	UserID    string `json:"user_id" query:"user-id"`
}

type ResponseGetListPost struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	UserID    string  `json:"user_id"`
	Agent     string  `json:"agent"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type (
	ResponseGetDetailPost struct {
		ID        string                   `json:"id"`
		Title     string                   `json:"title"`
		Body      string                   `json:"body"`
		UserID    string                   `json:"user_id"`
		Agent     string                   `json:"agent"`
		CreatedAt *string                  `json:"created_at"`
		UpdatedAt *string                  `json:"updated_at"`
		Comments  []ResponseGetListComment `json:"comments"`
	}

	ResponseGetListComment struct {
		ID        string  `json:"id"`
		Body      string  `json:"body"`
		UserID    string  `json:"user_id"`
		Agent     string  `json:"agent"`
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}
)

type ParamPost struct {
	ID string `json:"id" params:"id"`
}
