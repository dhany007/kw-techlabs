package controllers

import (
	"encoding/json"
	"latihan/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodPost {
		models.MethodNotAllowed(w)
		return
	}

	request := models.User{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := models.Response{
			Error: "BAD REQUEST",
		}
		models.ResponseJSON(w, response)
		return
	}

	user := &models.User{
		Username: request.Username,
		Password: request.Password,
	}

	models.AddUser(user)

	response := models.Response{
		Payload: user,
	}
	models.ResponseJSON(w, response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodGet {
		models.MethodNotAllowed(w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	users := models.GetUsers()

	response := models.Response{
		Payload: users,
	}

	models.ResponseJSON(w, response)
}
