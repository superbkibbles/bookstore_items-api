package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/superbkibbles/bookstore_items-api/src/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8003",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
