package app

import (
	"github.com/alvaro259818/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
