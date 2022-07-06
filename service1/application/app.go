package application

import (
	"context"
	"encoding/json"
	"golang_projects/iman-task/service1/repository"
	"net/http"
)

const (
	pages = 50
	url   = "https://gorest.co.in/public/v1/posts"
)

type App struct {
	postRepo repository.Repository
}

func New(postRepo repository.Repository) *App {
	return &App{
		postRepo: postRepo,
	}
}
func (a *App) DownloadPosts(ctx context.Context, arg interface{}) (interface{}, error) {
	url := arg.(string)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	var response []repository.Post
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, err
}
