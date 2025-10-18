package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {

	routes := httprouter.New()
	routes.HandlerFunc(http.MethodGet, "/health", HealthCheck)
	return routes
}
