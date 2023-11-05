package routers

import (
	"service1/controllers"

	"service1/db"

	"github.com/gorilla/mux"
)

func Route_func(router *mux.Router) {
	client := db.GetGlobalMongoClient()
	router.HandleFunc("/create_player", controllers.Create_player(client)).Methods("POST")
	router.HandleFunc("/read_all_players", controllers.Read_all_players(client)).Methods("GET")
	router.HandleFunc("/update_player/{PlayerID}", controllers.Update_player(client)).Methods("POST")
}
