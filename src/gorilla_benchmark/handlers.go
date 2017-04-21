package gorilla_benchmark

import (
	"encoding/json"
	"net/http"
	"strconv"

	"model"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ids := []int{}
	for id, _ := range model.Users {
		ids = append(ids, id)
	}

	json.NewEncoder(w).Encode(ids)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}

	json.NewDecoder(r.Body).Decode(u)

	model.InsertUser(u)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": u.Id()})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	u := get_context_user(r)

	if nil == u {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: "Not Found",
		}); err != nil {
			panic(err)
		}
		return
	}

	json.NewEncoder(w).Encode(u)
}

func ModifyUser(w http.ResponseWriter, r *http.Request) {
	u := get_context_user(r)

	if nil == u {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: "Not Found",
		}); err != nil {
			panic(err)
		}
		return
	}

	json.NewDecoder(r.Body).Decode(u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	u := get_context_user(r)

	if nil == u {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: "Not Found",
		}); err != nil {
			panic(err)
		}
		return
	}

	delete(model.Users, u.Id())
}

func get_context_user(r *http.Request) *model.User {
	vars := mux.Vars(r)
	var user_id int
	var err error
	if user_id, err = strconv.Atoi(vars["user_id"]); err != nil {
		panic(err)
	}
	if user, exists := model.Users[user_id]; exists {
		return user
	} else {
		return nil
	}
}

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		panic(err)
// 	}
// }

// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	var todoId int
// 	var err error
// 	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
// 		panic(err)
// 	}
// 	todo := RepoFindTodo(todoId)
// 	if todo.Id > 0 {
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(http.StatusOK)
// 		if err := json.NewEncoder(w).Encode(todo); err != nil {
// 			panic(err)
// 		}
// 		return
// 	}

// 	// If we didn't find it, 404
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusNotFound)
// 	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
// 		panic(err)
// 	}

// }

// /*
// Test with this curl command:
// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
// */
// func TodoCreate(w http.ResponseWriter, r *http.Request) {
// 	var todo Todo
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := r.Body.Close(); err != nil {
// 		panic(err)
// 	}
// 	if err := json.Unmarshal(body, &todo); err != nil {
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(422) // unprocessable entity
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			panic(err)
// 		}
// 	}

// 	t := RepoCreateTodo(todo)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusCreated)
// 	if err := json.NewEncoder(w).Encode(t); err != nil {
// 		panic(err)
// 	}
// }
