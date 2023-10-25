package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func reverseProxyHandler(targetURL string) http.Handler {
	target, _ := url.Parse(targetURL)
	return httputil.NewSingleHostReverseProxy(target)
}

func main() {
	router := mux.NewRouter()

	// Update the target URLs to connect to the services by their container names
	backendService1URL := "http://service1:8081" // Use the container name "service1"
	backendService2URL := "http://service2:8082" // Use the container name "service2"
	backendService3URL := "http://teams:8083"    // Use the container name "teams"
	// Create routes for the API gateway
	router.PathPrefix("/service1").Handler(reverseProxyHandler(backendService1URL))
	router.PathPrefix("/service2").Handler(reverseProxyHandler(backendService2URL))
	router.PathPrefix("/teams").Handler(reverseProxyHandler(backendService3URL))
	http.Handle("/", router)

	// Listen on port 8080 for external requests
	http.ListenAndServe(":8080", router)
}
