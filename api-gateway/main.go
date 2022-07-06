package main

import (
	"context"
	"log"

	"github.com/SherzodAbdullajonov/api-gateway/api"
	"github.com/SherzodAbdullajonov/api-gateway/database"
	"github.com/SherzodAbdullajonov/api-gateway/router"
	"github.com/SherzodAbdullajonov/api-gateway/server"
	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cfg, err := database.Load()
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}



	dataServiceClient, closeDataService, err := client.NewDataServiceClient(ctx, cfg.DataServiceAddr())
	if err != nil {
		log.Panicf("failed to connect to data service: %v", err)
	}
	defer closeDataService()

	postServiceClient, closePostService, err := client.NewPostServiceClient(ctx, cfg.PostServiceAddr())
	if err != nil {
		log.Panicf("failed to connect to post service: %v", err)
	}
	defer closePostService()

	server.RunHTTPServer(cfg.ListenAddress(), func(router mux.Router) http.Handler {
		api.InitRoutes(cfg, router, router.NewHTTPServer(dataServiceClient, postServiceClient))
		return router
	})
}
