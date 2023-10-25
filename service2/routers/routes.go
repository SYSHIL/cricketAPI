package routers

import (
	"service2/controllers"
	"service2/db"

	"github.com/gorilla/mux"
)

func Route_func(router *mux.Router) {
	client := db.GetGlobalMongoClient()
	router.HandleFunc("/delete_player/{PlayerID}", controllers.Delete_player(client)).Methods("DELETE")
}
