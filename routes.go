package main

import "net/http"

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"admin",
		[]string{"GET", "DELETE"},
		"/api/admin/{id}",
		AdminInfo,
	},
	Route{
		"admin",
		[]string{"POST"},
		"/api/admin",
		AdminInfo,
	},
	Route{
		"adminList",
		[]string{"GET"},
		"/api/adminlist",
		AdminsList,
	},
}
