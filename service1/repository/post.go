package repository

import "context"

type Post struct {
	Body   string `json:"body"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"user_id"`
}
type Pagination struct {
	Limit int64 `json:"limit"`
	Links struct {
	  Current  string      `json:"current"`
	  Next     string      `json:"next"`
	  Previous interface{} `json:"previous"`
	} `json:"links"`
	Page  int64 `json:"page"`
	Pages int64 `json:"pages"`
	Total int64 `json:"total"`
  } 

type Repository interface {
	CreatePosts(ctx context.Context)
}

func New (url string, limit int64) *Pagination{
	return &Pagination{
		Limit:  limit,
		Links: url,
		

	}
}