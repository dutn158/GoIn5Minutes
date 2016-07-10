package main

import "net/http"

// Route construct of a request
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HanderFunc http.HandlerFunc
}

// Routes slices of Route
type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
}
