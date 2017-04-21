package model

func GetUsers() []int {
	ids := []int{}
	for id := range Users {
		ids = append(ids, id)
	}

	return ids
}

func InsertUser(u *User) (error, interface{}) {
	users_last_id++ // NOTE: This should be thread safe in a nice server

	u.id = users_last_id
	Users[u.id] = u

	return nil, map[string]interface{}{"_id": u.id}
}
