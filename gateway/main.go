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
	BACKENDSERVICE1URL := os.getenv("BACKENDSERVICE1URL")
	BACKENDSERVICE1PORT := os.getenv("BACKENDSERVICE1PORT")

	BACKENDSERVICE2URL := os.getenv("BACKENDSERVICE2URL")
	BACKENDSERVICE2PORT := os.getenv("BACKENDSERVICE2PORT")

	BACKENDSERVICE3URL := os.getenv("BACKENDSERVICE3URL")
	BACKENDSERVICE3PORT := os.getenv("BACKENDSERVICE3PORT")

	GATEWAYPORT := os.getenv("GATEWAYPORT")

	backendService1URL := "http://" + BACKENDSERVICE1URL + ":" + BACKENDSERVICE1PORT // Use the container name "service1"
	backendService2URL := "http://" + BACKENDSERVICE2URL + ":" + BACKENDSERVICE2PORT // Use the container name "service2"
	backendService3URL := "http://" + BACKENDSERVICE3URL + ":" + BACKENDSERVICE3PORT // Use the container name "teams"

	// Create routes for the API gateway
	router.PathPrefix("/service1").Handler(reverseProxyHandler(backendService1URL))
	router.PathPrefix("/service2").Handler(reverseProxyHandler(backendService2URL))
	router.PathPrefix("/teams").Handler(reverseProxyHandler(backendService3URL))
	http.Handle("/", router)

	// Listen on port 8080 for external requests
	http.ListenAndServe(":"+GATEWAYPORT, router)
}
