package handler
import (
	"context"
	"golang_projects/iman-task/service2/repository"
)

type GetPostHandler struct {
	postRepo repository.Repository
}
type ListPostsHandler struct {
	postRepo repository.Repository
}

type UpdatePostHandler struct {
	postRepo repository.Repository
}
type DeletePostHandler struct {
	postRepo repository.Repository
}

func NewGetPostHandler(postRepo repository.Repository) GetPostHandler {
	return GetPostHandler{
		postRepo: postRepo,
	}
}
func NewListPostsHandler(postRepo repository.Repository) ListPostsHandler {
	return ListPostsHandler{
		postRepo: postRepo,
	}
}
func NewUpdatePostHandler(postRepo repository.Repository) UpdatePostHandler {
	return UpdatePostHandler{
		postRepo: postRepo,
	}
}
func NewDeletePostHandler(postRepo repository.Repository) DeletePostHandler {
	return DeletePostHandler{
		postRepo: postRepo,
	}
}

func (h GetPostHandler) GetPost(ctx context.Context, postID int) (repository.Post, error) {
	return h.postRepo.GetPost(ctx, postID)
}
func (h ListPostsHandler) GetPosts(ctx context.Context, page, limit int) ([]repository.Post, error) {
	if page < 1 {
		page = 1
	}
	if limit < 10 {
		limit = 10
	}
	
	return h.postRepo.ListPosts(ctx, page, limit)
}
func (h UpdatePostHandler) Handle(ctx context.Context, p repository.Post) error {
	return h.postRepo.UpdatePost(ctx, p)
}
func (h DeletePostHandler) Handle(ctx context.Context, postID int) error {
	return h.postRepo.DeletePost(ctx, postID)
}