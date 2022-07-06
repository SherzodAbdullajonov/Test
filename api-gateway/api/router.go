package api

import (
	"github.com/SherzodAbdullajonov/api-gateway/database"
	"github.com/SherzodAbdullajonov/api-gateway/repository"

	"github.com/gorilla/mux"
)

func InitRoutes(cfg database.Config, router mux.Router, s ports.HTTPServer) {

	router.Post("/posts/download", s.DownloadPosts)
	router.Get("/posts/download/status", s.GetDownloadStatus)
	router.Get("/posts/{post-id}", s.GetPost)
	router.Get("/posts", s.ListPosts)
	router.Put("/posts", s.UpdatePost)
	router.Delete("/posts/{post-id}", s.DeletePost)
}
