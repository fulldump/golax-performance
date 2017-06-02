package gorilla_benchmark

import (
	"encoding/json"
	"model"
	"net/http"
)

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

func init() {
	for _, n1 := range model.Letters {
		p := n1

		r := Route{
			"GetUser",
			"GET",
			"/letters/" + p,
			HandlerLetter(p),
		}

		routes = append(routes, r)

		for _, n2 := range model.Letters {
			p := n1 + "/" + n2

			r := Route{
				"GetUser",
				"GET",
				"/letters/" + p,
				HandlerLetter(p),
			}

			routes = append(routes, r)

			for _, n3 := range model.Letters {

				p := n1 + "/" + n2 + "/" + n3

				r := Route{
					"GetUser",
					"GET",
					"/letters/" + p,
					HandlerLetter(p),
				}

				routes = append(routes, r)
			}
		}
	}
}

func HandlerLetter(p string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"letter": "I am letter " + p,
		})
	}
}
