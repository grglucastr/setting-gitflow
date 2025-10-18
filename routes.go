package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	routes := httprouter.New()
	routes.HandlerFunc(http.MethodGet, "/health", app.healthCheck)
	return routes
}
