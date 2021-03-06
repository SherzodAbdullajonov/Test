package repository

import (
	"context"
	"errors"
)

type Post struct {
	id     int
	userID int
	title  string
	body   string
}

func New(id, userID int, title, body string) (Post, error) {
	p := Post{
		id:     id,
		userID: userID,
		title:  title,
		body:   body,
	}
	if err := p.validate(); err != nil {
		return Post{}, err
	}

	return p, nil
}

func (p Post) ID() int {
	return p.id
}

func (p Post) UserID() int {
	return p.userID
}

func (p Post) Title() string {
	return p.title
}

func (p Post) Body() string {
	return p.body
}

func (p Post) validate() error {
	if p.id == 0 {
		return errors.New("required id is not provided")
	}
	if p.userID < 0 {
		return errors.New("user id is negative")
	}
	if p.title == "" {
		return errors.New("required title is not provided")
	}
	if p.body == "" {
		return errors.New("required body is not provided")
	}

	return nil
}

type Repository interface {
	GetPost(ctx context.Context, postID int) (Post, error)
	ListPosts(ctx context.Context, page, limit int) ([]Post, error)
	UpdatePost(ctx context.Context, p Post) error
	DeletePost(ctx context.Context, postID int) error
}
