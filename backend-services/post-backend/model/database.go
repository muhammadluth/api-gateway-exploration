package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:public.post,alias:p_p"`
	ID            string     `json:"id" bun:",pk,unique,notnull,default:uuid_generate_v4()"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	UserID        string     `json:"user_id"`
	Agent         string     `json:"agent"`
	CreatedAt     time.Time  `json:"created_at" bun:",nullzero,notnull,default:now()"`
	UpdatedAt     time.Time  `json:"updated_at" bun:",nullzero,notnull,default:now()"`
	Comments      []*Comment `bun:"rel:has-many,join:id=post_id"`
}

type Comment struct {
	bun.BaseModel `bun:"table:public.comment,alias:p_c"`
	ID            string    `json:"id" bun:",pk,unique,notnull,default:uuid_generate_v4()"`
	Body          string    `json:"body"`
	UserID        string    `json:"user_id"`
	PostID        string    `json:"post_id"`
	Agent         string    `json:"agent"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:now()"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:now()"`
	Post          *Post     `bun:"rel:has-one,join:post_id=id"`
}
