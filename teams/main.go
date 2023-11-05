package main

import (
	"net/http"
	"teams/routers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a Gorilla Mux router
	router := mux.NewRouter()

	// Create a subrouter for the "/teams" path prefix
	subrouter := router.PathPrefix("/teams").Subrouter()

	// Set up routes using the subrouter
	routers.Route_func(subrouter) // Make sure this correctly sets up routes within "/service1"

	// Listen on port 8083 for all requests
	http.ListenAndServe(":"+BACKENDSERVICE3PORT, router)
}
