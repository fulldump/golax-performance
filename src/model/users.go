package model

type User struct {
	id           int    `json:"-"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Introduction string `json:"introduction"`
}

func (user *User) Id() int {
	return user.id
}

var Users = map[int]*User{}
var users_last_id = 0

/**
 * Insert 3 sample users
 */
func init() {
	InsertUser(&User{
		Name:         "Fulanito Fulanitez",
		Age:          20,
		Introduction: "Hello, I like flowers and plants",
	})

	InsertUser(&User{
		Name:         "Menganito Menganez",
		Age:          30,
		Introduction: "Hi, I like wheels and cars",
	})

	InsertUser(&User{
		Name:         "Zutanito Zutanez",
		Age:          40,
		Introduction: "Hey, I love cats and dogs",
	})
}
