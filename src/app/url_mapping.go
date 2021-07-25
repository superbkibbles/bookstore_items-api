package app

import (
	"net/http"

	"github.com/superbkibbles/bookstore_items-api/src/controllers"
)

func mapUrls() {
	// Ping controller
	router.HandleFunc("/ping", controllers.PingController.Ping)
	// Item controller
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemController.Search).Methods(http.MethodPost)
}
