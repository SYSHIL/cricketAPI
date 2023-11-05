package main

import (
	"net/http"
	"service2/routers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Create a subrouter for the "/service2" path prefix
	subrouter := router.PathPrefix("/service2").Subrouter()

	// Set up routes using the subrouter
	routers.Route_func(subrouter) // Make sure this correctly sets up routes within "/service2"

	// Listen on port 8082 for all requests
	http.ListenAndServe(":"+BACKENDSERVICE2PORT, router)

}
