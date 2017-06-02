package gorilla_benchmark

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
		"GetUsers",
		"GET",
		"/service/v1/users",
		GetUsers,
	},
	Route{
		"CreateUser",
		"POST",
		"/service/v1/users",
		CreateUser,
	},
	Route{
		"GetUser",
		"GET",
		"/service/v1/users/{user_id}",
		GetUser,
	},
	Route{
		"ModifyUser",
		"POST",
		"/service/v1/users/{user_id}",
		ModifyUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/service/v1/users/{user_id}",
		DeleteUser,
	},
}
