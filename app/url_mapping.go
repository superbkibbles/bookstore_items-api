package app

import (
	"net/http"

	"github.com/superbkibbles/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
}
