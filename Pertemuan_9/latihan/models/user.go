package models

var Users []User

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddUser(u *User) {
	Users = append(Users, *u)
}

func GetUsers() *[]User {
	return &Users
}
