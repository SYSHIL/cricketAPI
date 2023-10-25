package main

import (
	"net/http"
	"service1/routers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a Gorilla Mux router
	router := mux.NewRouter()

	// Create a subrouter for the "/service1" path prefix
	subrouter := router.PathPrefix("/service1").Subrouter()

	// Set up routes using the subrouter
	routers.Route_func(subrouter) // Make sure this correctly sets up routes within "/service1"

	// Listen on port 8081 for all requests
	http.ListenAndServe("0.0.0.0:8081", router)
}
