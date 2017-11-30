package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"admin",
		"GET",
		"/api/admin/{id}",
		AdminInfo,
	},
	Route{
		"admin",
		"POST",
		"/api/admin",
		AdminInfo,
	},
	Route{
		"admin",
		"DELETE",
		"/api/admin/{id}",
		AdminInfo,
	},
	Route{
		"AdminList",
		"GET",
		"/api/adminlist",
		AdminsList,
	},
}
