package chi_benchmark

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"strconv"

	"context"

	"github.com/pressly/chi"
)

func NewApi() http.Handler {

	r := chi.NewRouter()

	//r.Use(middleware.Logger)

	users := r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			ids := []int{}
			for id := range model.Users {
				ids = append(ids, id)
			}

			json.NewEncoder(w).Encode(ids)

		})

		r.Route("/:id", func(r chi.Router) {

			r.Use(UserCtx)

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				u := r.Context().Value("user")
				json.NewEncoder(w).Encode(u)
			})

		})

	})

	r.Mount("/service/v1/users", users)

	r.Route("/letters", func(r chi.Router) {
		MakeLetters(r, 3, "")
	})

	return r
}

func MakeLetters(r chi.Router, level int, prefix string) {

	if 0 == level {
		return
	}

	for _, letter := range model.Letters {
		p := prefix + "/" + letter
		MakeLetters(r.Route("/"+letter, func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"letter": "I am letter " + p,
				})
			})
		}), level-1, p)
	}

}

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user_id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if user, exists := model.Users[user_id]; exists {
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println("User", user_id, "not found 404")
			return
		}

	})
}
