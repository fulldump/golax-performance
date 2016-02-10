package main

import (
	"encoding/json"

	"model"

	"github.com/fulldump/golax"
)

func main() {

	my_api := golax.NewApi()
	my_api.Prefix = "/service/v1"

	my_api.Root.
		Interceptor(interceptor_logger).
		Interceptor(golax.InterceptorError)

	users := my_api.Root.Node("users").
		Method("GET", get_users).
		Method("POST", post_users)

	users.Node("{user_id}").
		Interceptor(interceptor_user).
		Method("GET", get_user).
		Method("POST", post_user).
		Method("DELETE", delete_user)

	my_api.Serve()
}

func get_users(c *golax.Context) {
	ids := []int{}
	for id, _ := range model.Users {
		ids = append(ids, id)
	}

	json.NewEncoder(c.Response).Encode(ids)
}

func post_users(c *golax.Context) {
	u := &model.User{}

	json.NewDecoder(c.Request.Body).Decode(u)

	model.InsertUser(u)

	c.Response.WriteHeader(201)
	json.NewEncoder(c.Response).Encode(map[string]interface{}{"id": u.Id()})
}

func get_user(c *golax.Context) {
	u := get_context_user(c)

	json.NewEncoder(c.Response).Encode(u)
}

func post_user(c *golax.Context) {
	u := get_context_user(c)

	json.NewDecoder(c.Request.Body).Decode(u)
}

func delete_user(c *golax.Context) {
	u := get_context_user(c)
	delete(model.Users, u.Id())
}

/**
 * Helper to get a user object from the context
 */
func get_context_user(c *golax.Context) *model.User {
	v, _ := c.Get("user")
	return v.(*model.User)
}
