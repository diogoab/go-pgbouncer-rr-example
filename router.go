package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all the routes
type Routes []Route

var routes = Routes{
	Route{
		"TodoCreate",
		"POST",
		"/todo",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todo/{todoID}",
		TodoGet,
	},
}

// NewRouter func returns a new router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
