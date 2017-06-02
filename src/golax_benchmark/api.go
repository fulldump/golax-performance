package golax_benchmark

import (
	"github.com/fulldump/golax"

	"encoding/json"
	"model"
	"net/http"
)

func NewApi() http.Handler {

	my_api := golax.NewApi()
	//my_api.Prefix = "/service/v1"

	root := my_api.Root.
		//		Interceptor(interceptor_logger).
		Interceptor(golax.InterceptorError)

	v1 := root.Node("service").Node("v1")

	users := v1.Node("users").
		Method("GET", get_users).
		Method("POST", post_users)

	users.Node("{user_id}").
		Interceptor(interceptor_user).
		Method("GET", get_user).
		Method("POST", post_user).
		Method("DELETE", delete_user)

	MakeLetters(my_api.Root.Node("letters"), 3, "")

	return my_api
}

func MakeLetters(r *golax.Node, level int, prefix string) {

	if 0 == level {
		return
	}

	for _, letter := range model.Letters {
		p := prefix + "/" + letter

		MakeLetters(r.Node(letter).Method("GET", func(c *golax.Context) {
			json.NewEncoder(c.Response).Encode(map[string]interface{}{
				"letter": "I am letter " + p,
			})
		}), level-1, p)
	}

}

func get_users(c *golax.Context) {

	json.NewEncoder(c.Response).Encode(model.GetUsers())
}

func post_users(c *golax.Context) {
	u := &model.User{}

	json.NewDecoder(c.Request.Body).Decode(u)

	e, r := model.InsertUser(u)
	if nil != e {
		// TODO: ERROR
		return
	}

	c.Response.WriteHeader(201)
	json.NewEncoder(c.Response).Encode(r)
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
