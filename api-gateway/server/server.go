package server

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunHTTPServer(addr string, createHandler func(router mux.Router) http.Handler){
	apiRouter := mux.NewRouter()

	rootRouter := mux.NewRouter()
	rootRouter.Handle("/api", createHandler(apiRouter))

	fmt.Println("Starting HTTP server")
	log.Fatal(http.ListenAndServe(addr, rootRouter))
}
