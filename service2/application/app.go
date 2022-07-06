package application

import (
	"github.com/SherzodAbdullajonov/service2/application/handler"
	"github.com/SherzodAbdullajonov/service2/repository"
)

type App struct {
	Handler Handler
}
type Handler struct {
	GetPost    handler.GetPostHandler
	ListPosts  handler.ListPostsHandler
	UpdatePost handler.UpdatePostHandler
	DeletePost handler.DeletePostHandler
}
func New (postRepo repository.Repository) App {
	return App{
		Handler: Handler{
			GetPost: handler.NewGetPostHandler(postRepo),
			ListPosts: handler.NewListPostsHandler(postRepo),
			UpdatePost: handler.NewUpdatePostHandler(postRepo),
			DeletePost: handler.NewDeletePostHandler(postRepo),
		},
	}
}
