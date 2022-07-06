package repository

import (
	"context"
	"fmt"
)

type Post struct {
	Body   string `json:"body"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"user_id"`
}
type Pagination struct {
	url         string
	limit       int64
	currentPage int
}

type Repository interface {
	CreatePosts(ctx context.Context)
}

func New(url string, limit int64) *Pagination {
	return &Pagination{
		limit:       limit,
		url:         url,
		currentPage: 0,
	}
}
func (u *Pagination) NextPage() (ok bool) {
	if u.currentPage >= int(u.limit) {
		return false
	}

	u.currentPage++
	return true
}
func (u *Pagination) URL() string {
	return fmt.Sprintf("%s?page=%d", u.url, u.currentPage)
}

func (u *Pagination) CurrentPage() int {
	return u.currentPage
}
