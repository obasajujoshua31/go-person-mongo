package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func LoadRoutes() []Route {
	return routes
}

func SetUpRoutes(r *mux.Router) *mux.Router {

	for _, route := range LoadRoutes() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
