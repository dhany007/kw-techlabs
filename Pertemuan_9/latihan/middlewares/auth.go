package middlewares

import (
	"latihan/models"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		response := models.Response{
			Error: "INVALID AUTH",
		}
		models.ResponseJSON(w, response)
		return false
	}

	isValid := false

	for _, v := range models.Users {
		if v.Username == username && v.Password == password {
			isValid = true
		}
	}

	return isValid
}
